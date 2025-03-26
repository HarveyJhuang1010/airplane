package seat

import (
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/domain/entities/po"
	"airplane/internal/enum"
	"airplane/internal/errs"
	"airplane/internal/tools/timelogger"
	"context"
	"github.com/pkg/errors"
)

func newBookSeat(in dependence) *BookSeat {
	return &BookSeat{
		in: in,
	}
}

type BookSeat struct {
	in dependence
}

func (uc *BookSeat) BookSeat(ctx context.Context, tx *rdb.Database, seatID int64) error {
	defer timelogger.LogTime(ctx)()
	seat, err := tx.SeatDAO().Get(ctx, seatID, true)
	if err != nil {
		uc.in.Logger.AppLogger.Error(ctx, err)
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrDBQueryFailed
		}
		return err
	}
	if seat.Status != enum.SeatStatusAvailable {
		return errs.ErrSeatNotAvailable
	} else {
		seat.Status = enum.SeatStatusBooked
		if err := tx.SeatDAO().Update(ctx, &po.SeatUpdateCond{
			ID:     seatID,
			Status: enum.SeatStatusBooked,
		}); err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return err
		}
	}
	return nil
}
