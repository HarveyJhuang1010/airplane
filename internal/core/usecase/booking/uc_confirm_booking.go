package booking

import (
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/domain/entities/bo"
	"airplane/internal/domain/entities/po"
	"airplane/internal/enum"
	"airplane/internal/errs"
	"airplane/internal/tools/timelogger"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
)

func newConfirmBooking(in digIn) *ConfirmBooking {
	return &ConfirmBooking{
		in: in,
	}
}

type ConfirmBooking struct {
	in digIn
}

func (uc *ConfirmBooking) ConfirmBooking(ctx context.Context, data []byte) error {
	defer timelogger.LogTime(ctx)()

	cond := &bo.ConfirmBookingCond{}
	if err := json.Unmarshal(data, cond); err != nil {
		uc.in.Logger.AppLogger.Error(ctx, err)
		return errs.ErrParseFailed.TraceWrap(err)
	}

	return uc.in.DBRepository.Master().Transaction(func(tx *rdb.Database) error {

		// get booking
		booking, err := tx.BookingDAO().Get(ctx, cond.ID, true, false)
		if err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return err
		}

		// check the status
		if booking.Status != enum.BookingStatusPending {
			return errs.ErrStatusNotMatch
		}

		// update the remain seat
		class, err := tx.CabinClassDAO().Get(ctx, booking.CabinClassID, true)
		if err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			if errors.Is(err, errs.ErrRecordNotFound) {
				return errs.ErrDBQueryFailed
			}
			return err
		}
		if class.RemainSeats <= 0 {
			booking.Status = enum.BookingStatusOverbooked
		} else {
			booking.Status = enum.BookingStatusConfirmed
			class.RemainSeats--
		}

		if err := tx.CabinClassDAO().Update(ctx, class.ID, class.RemainSeats); err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return err
		}

		// select seat(optional)
		if booking.SeatID != nil {
			if err := uc.in.Seat.BookSeat.BookSeat(ctx, tx, *booking.SeatID); err != nil {
				if errors.Is(err, errs.ErrSeatNotAvailable) {
					booking.Status = enum.BookingStatusOverbooked
					if err := tx.BookingDAO().CancelSeat(ctx, booking.ID); err != nil {
						uc.in.Logger.AppLogger.Error(ctx, err)
						return err
					}
				} else {
					return err
				}
			}
		}

		// update booking
		if err := tx.BookingDAO().UpdateStatus(ctx, &po.BookingUpdateCond{
			ID:     booking.ID,
			Status: booking.Status,
		}); err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return err
		}

		return nil
	})
}
