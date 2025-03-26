package rest

import (
	"airplane/internal/components/apis"
	"airplane/internal/components/logger"
	"airplane/internal/controller/portal/rest/bookingpkg"
	"airplane/internal/controller/portal/rest/flight"
	"airplane/internal/controller/portal/rest/payment"
	"go.uber.org/dig"
)

func NewBookingRouter(in digIn) {
	newBookingRouter(newRouter(in))
}

func NewFlightRouter(in digIn) {
	newFlightRouter(newRouter(in))
}

func NewPaymentRouter(in digIn) {
	newPaymentRouter(newRouter(in))
}

func newRouter(in digIn) dependence {
	dep := dependence{
		digIn: in,
	}

	return dep
}

type dependence struct {
	digIn
}

type digIn struct {
	dig.In

	Logger     *logger.Loggers
	ApiService *apis.Service `name:"portal_api_service"`

	PaymentCtrl *payment.Controller
	BookingCtrl *bookingpkg.Controller
	FlightCtrl  *flight.Controller
}
