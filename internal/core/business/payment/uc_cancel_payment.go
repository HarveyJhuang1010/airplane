package payment

import (
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/tools/timelogger"
	"context"
)

func newCancelPayment(in dependence) *CancelPayment {
	return &CancelPayment{
		in: in,
	}
}

type CancelPayment struct {
	in dependence
}

func (uc *CancelPayment) CancelPayment(ctx context.Context, tx *rdb.Database, bookingID int64) error {
	defer timelogger.LogTime(ctx)()
	// implement me
	return nil
}
