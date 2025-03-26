package payment

import (
	"airplane/internal/infrastructure/mqcli"
	"airplane/internal/infrastructure/snowflake"
	"go.uber.org/dig"
)

func New(in digIn) digOut {
	dep := dependence{
		digIn: in,
	}

	return digOut{
		Usecase: newPaymentUsecase(dep),
	}
}

type dependence struct {
	digIn
}

type digIn struct {
	dig.In

	Snowflake *snowflake.Snowflake
	Kafka     *mqcli.Kafka
}

type digOut struct {
	dig.Out

	Usecase *Usecase
}

func newPaymentUsecase(in dependence) *Usecase {
	return &Usecase{
		CreatePayment:  newCreatePayment(in),
		ConfirmPayment: newConfirmPayment(in),
		CancelPayment:  newCancelPayment(in),
		RefundPayment:  newRefundPayment(in),
	}
}

type Usecase struct {
	CreatePayment  *CreatePayment
	ConfirmPayment *ConfirmPayment
	CancelPayment  *CancelPayment
	RefundPayment  *RefundPayment
}
