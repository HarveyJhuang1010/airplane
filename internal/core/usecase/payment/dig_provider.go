package payment

import (
	"airplane/internal/components/logger"
	"airplane/internal/core/business/payment"
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/infrastructure/mqcli"
	"go.uber.org/dig"
)

func New(in digIn) digOut {
	dep := dependence{
		digIn: in,
	}

	return digOut{
		Usecase: newFlightUsecase(dep),
	}
}

type dependence struct {
	digIn
}

type digIn struct {
	dig.In

	Logger          *logger.Loggers
	DBRepository    *rdb.Repository
	Kafka           *mqcli.Kafka
	PaymentBusiness *payment.Usecase
}

type digOut struct {
	dig.Out

	Usecase *Usecase
}

func newFlightUsecase(in dependence) *Usecase {
	return &Usecase{
		PaymentWebhook: newPaymentWebhook(in),
	}
}

type Usecase struct {
	PaymentWebhook *PaymentWebhook
}
