package payment

import (
	"airplane/internal/components/apis"
	"airplane/internal/domain/entities/bo"
	"airplane/internal/tools/timelogger"
	"github.com/gin-gonic/gin"
)

func newPayment(in dependence) *Payment {
	return &Payment{
		in:       in,
		Response: in.Response,
	}
}

type Payment struct {
	in       dependence
	Response apis.IResponse
}

// NotifyPaymentResult
// @Summary  Notify Payment Result
// @Description 3rd party payment gateway will call this API to notify payment result
// @Tags 		Payment
// @Accept 		json
// @Produce		json
// @Param 		params body NotifyPaymentCond true "Request Body"
// @Success 	201 {object} apis.StandardResponse{}
// @Failure 	400 {object} apis.StandardResponse{error=apis.StandardError}
// @Failure 	500 {object} apis.StandardResponse{error=apis.StandardError}
// @Router 		/payment/notify [post]
func (ctrl *Payment) NotifyPaymentResult(ctx *gin.Context) {
	var err error
	defer func() {
		timelogger.LogTime(ctx)()
		if err != nil {
			ctrl.Response.Data(ctx, ctrl.in.StandardData().BadRequest(nil, err))
		}
	}()

	dto, err := apis.RequestParser[NotifyPaymentCond]().Json().Bind(ctx)
	if err != nil {
		return
	}

	cond, err := apis.Convert[bo.ConfirmPaymentCond](ctx, dto)
	if err != nil {
		return
	}

	if err := ctrl.in.Payment.PaymentWebhook.PaymentWebhook(ctx, cond); err != nil {
		return
	}

	ctrl.Response.Data(ctx, ctrl.in.StandardData().Created(nil))
}
