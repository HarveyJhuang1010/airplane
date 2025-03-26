package booking

import (
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/domain/entities/bo"
	"airplane/internal/domain/entities/po"
	"airplane/internal/enum"
	"airplane/internal/errs"
	"airplane/internal/tools/timelogger"
	"context"
	"time"
)

func newEditBooking(in digIn) *EditBooking {
	return &EditBooking{
		in: in,
	}
}

type EditBooking struct {
	in digIn
}

func (uc *EditBooking) EditBooking(ctx context.Context, cond *bo.EditBookingCond) error {
	defer timelogger.LogTime(ctx)()

	if cond == nil || cond.ID == 0 || cond.CabinClassID == 0 || cond.SeatID == 0 {
		return errs.ErrInvalidParameter
	}

	return uc.in.DBRepository.Master().Transaction(func(tx *rdb.Database) error {
		booking, err := tx.BookingDAO().Get(ctx, cond.ID, true, false)
		if err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return err
		}

		if booking.Status != enum.BookingStatusConfirmed && booking.Status != enum.BookingStatusOverbooked {
			return errs.ErrStatusNotMatch
		}

		if booking.SeatID == &cond.SeatID {
			return nil
		}

		newSeat, err := tx.SeatDAO().Get(ctx, cond.SeatID, true)
		if err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return err
		}

		if newSeat.Status != enum.SeatStatusAvailable {
			return errs.ErrStatusNotMatch
		}

		newClass, err := tx.CabinClassDAO().Get(ctx, cond.CabinClassID, true)
		if err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return err
		}

		if newSeat.CabinClassID != newClass.ID {
			return errs.ErrInvalidParameter
		}

		if !booking.Price.Equal(newClass.Price) {
			payment, err := tx.PaymentDAO().GetByBookingID(ctx, booking.ID, false)
			if err != nil {
				uc.in.Logger.AppLogger.Error(ctx, err)
				return err
			}

			payment.Amount = newClass.Price.Sub(payment.Amount)
			extraPayment := &po.ExtraPayment{
				ID:            uc.in.Snowflake.Generate().Int64(),
				BookingID:     booking.ID,
				UserID:        booking.UserID,
				PaymentStatus: enum.PaymentStatusPending,
				Amount:        newClass.Price.Sub(payment.Amount),
				ExpiredAt:     time.Now().Add(time.Hour),
			}

			if err := tx.ExtraPaymentDAO().Create(ctx, extraPayment); err != nil {
				uc.in.Logger.AppLogger.Error(ctx, err)
				return err
			}
			if err := tx.SeatDAO().Update(ctx, &po.SeatUpdateCond{
				ID:     newSeat.ID,
				Status: enum.SeatStatusHeld,
			}); err != nil {
				uc.in.Logger.AppLogger.Error(ctx, err)
				return err
			}
			booking.Status = enum.BookingStatusConfirming
		} else {
			if err := tx.SeatDAO().Update(ctx, &po.SeatUpdateCond{
				ID:     newSeat.ID,
				Status: enum.SeatStatusBooked,
			}); err != nil {
				uc.in.Logger.AppLogger.Error(ctx, err)
				return err
			}
			booking.Status = enum.BookingStatusConfirmed
		}

		booking.SeatID = &cond.SeatID
		booking.CabinClassID = cond.CabinClassID
		booking.Price = newClass.Price

		if err := tx.BookingDAO().UpdateSeat(ctx, &po.BookingUpdateSeatCond{
			ID:           booking.ID,
			SeatID:       &newSeat.ID,
			CabinClassID: &newClass.ID,
			Price:        &newClass.Price,
		}); err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return err
		}

		return nil
	})
}
