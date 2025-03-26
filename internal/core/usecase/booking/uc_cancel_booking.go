package booking

import (
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/domain/entities/po"
	"airplane/internal/enum"
	"airplane/internal/errs"
	"airplane/internal/tools/timelogger"
	"context"
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

		if booking.Status == enum.BookingStatusCancelled || booking.Status == enum.BookingStatusExpired {
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

		if booking.Status == enum.BookingStatusConfirmed || booking.Status == enum.BookingStatusOverbooked {
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
					Status: enum.SeatStatusAvailable,
				}); err != nil {
					uc.in.Logger.AppLogger.Error(ctx, err)
					return err
				}
			}
		} else if booking.Status == enum.BookingStatusPending {
			// cancel the payment
			if err := uc.in.Payment.CancelPayment.CancelPayment(ctx, tx, booking.ID); err != nil {
				uc.in.Logger.AppLogger.Error(ctx, err)
				return err
			}
		}

		return nil
	})
}
