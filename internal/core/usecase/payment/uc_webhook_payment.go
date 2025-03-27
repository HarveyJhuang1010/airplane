package payment

import (
	"airplane/internal/constant"
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/domain/entities/bo"
	"airplane/internal/errs"
	"airplane/internal/tools/timelogger"
	"context"
	"encoding/json"
)

func newPaymentWebhook(in dependence) *PaymentWebhook {
	return &PaymentWebhook{
		in: in,
	}
}

type PaymentWebhook struct {
	in dependence
}

func (uc *PaymentWebhook) PaymentWebhook(ctx context.Context, cond *bo.ConfirmPaymentCond) error {
	defer timelogger.LogTime(ctx)()

	return uc.in.DBRepository.Master().Transaction(func(tx *rdb.Database) error {
		// get payment
		payment, err := tx.PaymentDAO().Get(ctx, cond.ID, true)
		if err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return err
		}

		if err := uc.in.PaymentBusiness.ConfirmPayment.ConfirmPayment(ctx, tx, cond); err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return err
		}

		msgCond := &bo.ConfirmBookingCond{
			ID: payment.BookingID,
		}

		msgVal, err := json.Marshal(msgCond)
		if err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return errs.ErrParseFailed.TraceWrap(err)
		}

		if err := uc.in.Kafka.Produce(ctx, constant.KafkaTopicConfirmBooking, msgVal); err != nil {
			return errs.ErrMQFailed.TraceWrap(err)
		}

		return nil
	})
}
