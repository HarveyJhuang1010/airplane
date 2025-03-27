package payment

import (
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/domain/entities/bo"
	"airplane/internal/domain/entities/po"
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

	return tx.PaymentDAO().UpdateResult(ctx, &po.PaymentUpdateResultCond{
		ID:            cond.ID,
		TransactionID: cond.TransactionID,
		Provider:      cond.Provider,
		Method:        cond.Method,
		Status:        cond.Status,
		PaidAt:        cond.PaidAt,
	})
}
