package booking

import (
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/domain/entities/bo"
	"airplane/internal/domain/entities/po"
	"airplane/internal/enum"
	"airplane/internal/tools/timelogger"
	"context"
)

func newHandleOverBooking(in digIn) *HandleOverBooking {
	return &HandleOverBooking{
		in: in,
	}
}

type HandleOverBooking struct {
	in digIn
}

func (uc *HandleOverBooking) OverBookingToConfirm(ctx context.Context, tx *rdb.Database, flightID int64) (*bo.Booking, error) {
	defer timelogger.LogTime(ctx)()

	// release the overbooking
	var booking *po.Booking
	overbookings, err := tx.BookingDAO().GetOverBooking(ctx, flightID)
	if err != nil {
		uc.in.Logger.AppLogger.Error(ctx, err)
		return nil, err
	}

	if len(overbookings) > 0 {
		booking = overbookings[0]
		if err := tx.BookingDAO().UpdateStatus(ctx, &po.BookingUpdateCond{
			ID:     booking.ID,
			Status: enum.BookingStatusConfirmed,
		}); err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return nil, err
		}
		return &bo.Booking{
			ID:           booking.ID,
			FlightID:     booking.FlightID,
			UserID:       booking.UserID,
			CabinClassID: booking.CabinClassID,
			SeatID:       booking.SeatID,
			Status:       enum.BookingStatusConfirmed,
			Price:        booking.Price,
			ExpiredAt:    booking.ExpiredAt,
		}, nil
	}

	return nil, nil
}
