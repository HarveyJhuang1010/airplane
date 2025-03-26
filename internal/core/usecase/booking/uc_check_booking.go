package booking

import (
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/domain/entities/po"
	"airplane/internal/enum"
	"context"
	"time"

	"airplane/internal/tools/timelogger"
)

func newCheckBooking(in digIn) *CheckBooking {
	return &CheckBooking{
		in: in,
	}
}

type CheckBooking struct {
	in digIn
}

func (uc *CheckBooking) CheckBooking(ctx context.Context) error {
	defer timelogger.LogTime(ctx)()
	now := time.Now()

	// get expired booking
	expired, err := uc.in.DBRepository.Master().BookingDAO().GetExpired(ctx, now)
	if err != nil {
		uc.in.Logger.AppLogger.Error(ctx, err)
		return err
	}

	if len(expired) == 0 {
		return nil
	}

	return uc.in.DBRepository.Master().Transaction(func(tx *rdb.Database) error {
		for _, booking := range expired {
			// update the booking status
			if err := tx.BookingDAO().UpdateStatus(ctx, &po.BookingUpdateCond{
				ID:     booking.ID,
				Status: enum.BookingStatusExpired,
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

			// cancel the payment
			if err := uc.in.Payment.CancelPayment.CancelPayment(ctx, tx, booking.ID); err != nil {
				uc.in.Logger.AppLogger.Error(ctx, err)
				return err
			}
		}

		return nil
	})
}
