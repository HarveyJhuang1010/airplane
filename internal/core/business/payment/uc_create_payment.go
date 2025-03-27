package payment

import (
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/domain/entities/bo"
	"airplane/internal/domain/entities/po"
	"airplane/internal/enum"
	"airplane/internal/tools/timelogger"
	"context"
	"fmt"
	"time"
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

	id := uc.in.Snowflake.Generate().Int64()
	payment := &po.Payment{
		ID:         id,
		BookingID:  cond.BookingID,
		UserID:     cond.UserID,
		Provider:   nil,
		Method:     nil,
		Status:     enum.PaymentStatusPending,
		Amount:     cond.Amount,
		PaymentURL: fmt.Sprintf("https://mockpayment.com/%d", id),
		ExpiredAt:  time.Now().AddDate(0, 0, 1),
	}
	if err := tx.PaymentDAO().Create(ctx, payment); err != nil {
		return nil, err
	}

	return &bo.Payment{
		ID:            payment.ID,
		BookingID:     payment.BookingID,
		Amount:        payment.Amount,
		PaymentStatus: payment.Status,
		PaymentURL:    payment.PaymentURL,
	}, nil
}
