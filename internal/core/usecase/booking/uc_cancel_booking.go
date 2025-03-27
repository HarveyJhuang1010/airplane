package booking

import (
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/domain/entities/po"
	"airplane/internal/enum"
	"airplane/internal/errs"
	"airplane/internal/tools/timelogger"
	"context"
	"github.com/samber/lo"
)

func newCancelBooking(in digIn) *CancelBooking {
	return &CancelBooking{
		in: in,
	}
}

type CancelBooking struct {
	in digIn
}

func (uc *CancelBooking) CancelBooking(ctx context.Context, id int64) error {
	defer timelogger.LogTime(ctx)()

	return uc.in.DBRepository.Master().Transaction(func(tx *rdb.Database) error {
		booking, err := tx.BookingDAO().Get(ctx, id, true, false)
		if err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return err
		}

		if lo.Contains([]enum.BookingStatus{
			enum.BookingStatusConfirming,
			enum.BookingStatusCancelled,
			enum.BookingStatusExpired,
		}, booking.Status) {
			return errs.ErrStatusNotMatch
		}

		// update the booking status
		if err := tx.BookingDAO().UpdateStatus(ctx, &po.BookingUpdateCond{
			ID:     booking.ID,
			Status: enum.BookingStatusCancelled,
		}); err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return err
		}

		// update flight sellable seats
		flight, err := tx.FlightDAO().Get(ctx, booking.FlightID, true)
		if err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return err
		}
		flight.SellableSeats++

		if err := tx.FlightDAO().UpdateSellableSeats(ctx, flight.ID, flight.SellableSeats); err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return err
		}

		if lo.Contains([]enum.BookingStatus{enum.BookingStatusConfirmed, enum.BookingStatusOverbooked}, booking.Status) {
			// refund the payment
			if err := uc.in.Payment.RefundPayment.RefundPayment(ctx, tx, booking.ID); err != nil {
				uc.in.Logger.AppLogger.Error(ctx, err)
				return err
			}
			// update the remain seat
			class, err := tx.CabinClassDAO().Get(ctx, booking.CabinClassID, true)
			if err != nil {
				uc.in.Logger.AppLogger.Error(ctx, err)
				return err
			}
			class.RemainSeats++

			if err := tx.CabinClassDAO().Update(ctx, class.ID, class.RemainSeats); err != nil {
				uc.in.Logger.AppLogger.Error(ctx, err)
				return err
			}

			// release the overbooking
			if class.RemainSeats > 0 {
				_, err := self.Usecase.HandleOverBooking.OverBookingToConfirm(ctx, tx, booking.FlightID)
				if err != nil {
					uc.in.Logger.AppLogger.Error(ctx, err)
					return err
				}
			}

			// update the seat status
			if booking.SeatID != nil {
				if err := tx.SeatDAO().Update(ctx, &po.SeatUpdateCond{
					ID:     *booking.SeatID,
					Status: enum.SeatStatusHeld,
				}); err != nil {
					uc.in.Logger.AppLogger.Error(ctx, err)
					return err
				}
			}
		} else {
			// cancel the payment
			if err := uc.in.Payment.CancelPayment.CancelPayment(ctx, tx, booking.ID); err != nil {
				uc.in.Logger.AppLogger.Error(ctx, err)
				return err
			}
		}

		return nil
	})
}
