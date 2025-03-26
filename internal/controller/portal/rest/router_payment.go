package rest

func newPaymentRouter(dependence dependence) *PaymentRouter {
	return (&PaymentRouter{
		in: dependence,
	}).New()
}

type PaymentRouter struct {
	in dependence
}

func (r *PaymentRouter) New() *PaymentRouter {
	group := r.in.ApiService.GetStandardRouterGroup()

	// 前綴
	paymentPortalGroup := group.Group("/payment")

	paymentPortalGroup.POST("/notify", r.in.PaymentCtrl.Payment.NotifyPaymentResult)

	return r
}
