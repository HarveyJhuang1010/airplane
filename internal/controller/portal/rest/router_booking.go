package rest

func newBookingRouter(dependence dependence) *BookingRouter {
	return (&BookingRouter{
		in: dependence,
	}).New()
}

type BookingRouter struct {
	in dependence
}

func (r *BookingRouter) New() *BookingRouter {
	group := r.in.ApiService.GetStandardRouterGroup()

	bookingPortalGroup := group.Group("/booking")
	bookingPortalGroup.POST("", r.in.BookingCtrl.Booking.AddBooking)
	bookingPortalGroup.GET("/:id", r.in.BookingCtrl.Booking.GetBooking)
	bookingPortalGroup.DELETE("/:id", r.in.BookingCtrl.Booking.CancelBooking)
	bookingPortalGroup.PATCH("/:id", r.in.BookingCtrl.Booking.EditBooking)

	return r
}
