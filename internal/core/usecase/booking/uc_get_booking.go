package booking

import (
	"airplane/internal/domain/entities/bo"
	"airplane/internal/tools/timelogger"
	"context"
	"github.com/jinzhu/copier"
)

func newGetBooking(in digIn) *GetBooking {
	return &GetBooking{
		in: in,
	}
}

type GetBooking struct {
	in digIn
}

func (uc *GetBooking) GetBooking(ctx context.Context, id int64) (*bo.Booking, error) {
	defer timelogger.LogTime(ctx)()

	booking, err := uc.in.DBRepository.Master().BookingDAO().Get(ctx, id, false, true)
	if err != nil {
		uc.in.Logger.AppLogger.Error(ctx, err)
		return nil, err
	}

	result := &bo.Booking{}
	if err := copier.Copy(&result, booking); err != nil {
		uc.in.Logger.AppLogger.Error(ctx, err)
		return nil, err
	}

	return result, nil
}
