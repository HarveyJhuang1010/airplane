package payment

import (
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/domain/entities/bo"
	"airplane/internal/enum"
	"airplane/internal/tools/timelogger"
	"context"
)

func newCreatePayment(in dependence) *CreatePayment {
	return &CreatePayment{
		in: in,
	}
}

type CreatePayment struct {
	in dependence
}

func (uc *CreatePayment) CreatePayment(ctx context.Context, tx *rdb.Database, cond *bo.CreatePaymentCond) (*bo.Payment, error) {
	defer timelogger.LogTime(ctx)()
	// Implement the business logic of CreatePayment here
	return &bo.Payment{
		ID:            uc.in.Snowflake.Generate().Int64(),
		BookingID:     cond.BookingID,
		Amount:        cond.Amount,
		PaymentStatus: enum.PaymentStatusPending,
		PaymentURL:    "https://payment.com",
	}, nil
}
