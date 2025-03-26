package payment

import (
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/domain/entities/bo"
	"airplane/internal/tools/timelogger"
	"context"
)

func newConfirmPayment(in dependence) *ConfirmPayment {
	return &ConfirmPayment{
		in: in,
	}
}

type ConfirmPayment struct {
	in dependence
}

func (uc *ConfirmPayment) ConfirmPayment(ctx context.Context, tx *rdb.Database, cond *bo.ConfirmPaymentCond) error {
	defer timelogger.LogTime(ctx)()
	// Implement the business logic of NotifyPaymentResult here
	return nil
}
