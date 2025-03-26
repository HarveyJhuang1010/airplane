package rest

func newFlightRouter(dependence dependence) *FlightRouter {
	return (&FlightRouter{
		in: dependence,
	}).New()
}

type FlightRouter struct {
	in dependence
}

func (r *FlightRouter) New() *FlightRouter {
	group := r.in.ApiService.GetStandardRouterGroup()

	flightPortalGroup := group.Group("/api/v1/flight")
	flightPortalGroup.GET("", r.in.FlightCtrl.Flight.ListFlight)

	return r
}
