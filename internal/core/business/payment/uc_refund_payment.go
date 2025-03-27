package payment

import (
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/enum"
	"airplane/internal/errs"
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

	payment, err := tx.PaymentDAO().GetByBookingID(ctx, bookingID, true)
	if err != nil {
		return err
	}

	if payment.Status != enum.PaymentStatusSuccess {
		return errs.ErrStatusNotMatch
	}

	// refund the payment
	// implement me

	// update the payment status
	return tx.PaymentDAO().UpdateStatus(ctx, payment.ID, enum.PaymentStatusRefunding)
}
