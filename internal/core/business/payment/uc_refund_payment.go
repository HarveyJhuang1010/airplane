package payment

import (
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/tools/timelogger"
	"context"
)

func newRefundPayment(in dependence) *RefundPayment {
	return &RefundPayment{
		in: in,
	}
}

type RefundPayment struct {
	in dependence
}

func (uc *RefundPayment) RefundPayment(ctx context.Context, tx *rdb.Database, bookingID int64) error {
	defer timelogger.LogTime(ctx)()
	// implement me
	return nil
}
