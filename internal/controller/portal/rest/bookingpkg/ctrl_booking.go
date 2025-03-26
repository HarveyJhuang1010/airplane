package bookingpkg

import (
	"airplane/internal/components/apis"
	"airplane/internal/domain/entities/bo"
	"airplane/internal/tools/timelogger"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

func newBooking(in dependence) *Booking {
	return &Booking{
		in:       in,
		Response: in.Response,
	}
}

type Booking struct {
	in       dependence
	Response apis.IResponse
}

// AddBooking
// @Summary user booking ticket
// @Description user booking ticket
// @Tags 		Booking
// @Accept 		json
// @Produce		json
// @Param 		params body AddBookingCond true "Request Body"
// @Success 	200 {object} apis.StandardResponse{data=BookingResponse}
// @Failure 	400 {object} apis.StandardResponse{error=apis.StandardError}
// @Failure 	500 {object} apis.StandardResponse{error=apis.StandardError}
// @Router 		/booking [post]
func (ctrl *Booking) AddBooking(ctx *gin.Context) {
	var err error
	defer func() {
		timelogger.LogTime(ctx)()
		if err != nil {
			ctrl.Response.Data(ctx, ctrl.in.StandardData().BadRequest(nil, err))
		}
	}()

	dto, err := apis.RequestParser[AddBookingCond]().Json().Bind(ctx)
	if err != nil {
		return
	}

	cond, err := apis.Convert[bo.AddBookingCond](ctx, dto)
	if err != nil {
		return
	}

	if err := ctrl.in.Booking.AddBooking.AddBooking(ctx, cond); err != nil {
		return
	}

	booking, err := ctrl.in.Booking.GetBooking.GetBooking(ctx, cond.ID)
	if err != nil {
		return
	}

	ctrl.Response.Data(ctx, ctrl.in.StandardData().OK(&BookingResponse{
		ID:               booking.ID,
		AirlineCode:      booking.Flight.AirlineCode,
		FlightNumber:     booking.Flight.FlightNumber,
		DepartureAirport: booking.Flight.DepartureAirport,
		ArrivalAirport:   booking.Flight.ArrivalAirport,
		DepartureTime:    booking.Flight.DepartureTime,
		ArrivalTime:      booking.Flight.ArrivalTime,
		Email:            booking.User.Email,
		PhoneCountryCode: booking.User.PhoneCountryCode,
		PhoneNumber:      booking.User.PhoneNumber,
		ClassCode:        booking.Class.ClassCode,
		BaggageAllowance: booking.Class.BaggageAllowance,
		Refundable:       booking.Class.Refundable,
		SeatSelection:    booking.Class.SeatSelection,
		SeatNumber:       lo.Ternary(lo.IsNil(booking.Seat), booking.Seat.SeatNumber, ""),
		Status:           booking.Status,
		Price:            booking.Price,
		ExpiredAt:        booking.ExpiredAt,
	}))
}

// GetBooking
// @Summary user get booking ticket
// @Description user get booking ticket
// @Tags 		Booking
// @Accept 		json
// @Produce		json
// @Param		id path string true "id"
// @Success 	200 {object} apis.StandardResponse{data=BookingResponse}
// @Failure 	400 {object} apis.StandardResponse{error=apis.StandardError}
// @Failure 	500 {object} apis.StandardResponse{error=apis.StandardError}
// @Router 		/booking/:id [get]
func (ctrl *Booking) GetBooking(ctx *gin.Context) {
	var err error
	defer func() {
		timelogger.LogTime(ctx)()
		if err != nil {
			ctrl.Response.Data(ctx, ctrl.in.StandardData().BadRequest(nil, err))
		}
	}()

	dto, err := apis.RequestParser[GetBookingCond]().Uri().Bind(ctx)
	if err != nil {
		return
	}

	booking, err := ctrl.in.Booking.GetBooking.GetBooking(ctx, dto.ID)
	if err != nil {
		return
	}

	ctrl.Response.Data(ctx, ctrl.in.StandardData().OK(&BookingResponse{
		ID:               booking.ID,
		AirlineCode:      booking.Flight.AirlineCode,
		FlightNumber:     booking.Flight.FlightNumber,
		DepartureAirport: booking.Flight.DepartureAirport,
		ArrivalAirport:   booking.Flight.ArrivalAirport,
		DepartureTime:    booking.Flight.DepartureTime,
		ArrivalTime:      booking.Flight.ArrivalTime,
		Email:            booking.User.Email,
		PhoneCountryCode: booking.User.PhoneCountryCode,
		PhoneNumber:      booking.User.PhoneNumber,
		ClassCode:        booking.Class.ClassCode,
		BaggageAllowance: booking.Class.BaggageAllowance,
		Refundable:       booking.Class.Refundable,
		SeatSelection:    booking.Class.SeatSelection,
		SeatNumber:       lo.Ternary(lo.IsNil(booking.Seat), booking.Seat.SeatNumber, ""),
		Status:           booking.Status,
		Price:            booking.Price,
		ExpiredAt:        booking.ExpiredAt,
	}))
}

// CancelBooking
// @Summary user cancel booking ticket
// @Description user cancel booking ticket
// @Tags 		Booking
// @Accept 		json
// @Produce		json
// @Param		id path string true "id"
// @Success 	204
// @Failure 	400 {object} apis.StandardResponse{error=apis.StandardError}
// @Failure 	500 {object} apis.StandardResponse{error=apis.StandardError}
// @Router 		/booking/:id [delete]
func (ctrl *Booking) CancelBooking(ctx *gin.Context) {
	var err error
	defer func() {
		timelogger.LogTime(ctx)()
		if err != nil {
			ctrl.Response.Data(ctx, ctrl.in.StandardData().BadRequest(nil, err))
		}
	}()

	dto, err := apis.RequestParser[CancelBookingCond]().Uri().Bind(ctx)
	if err != nil {
		return
	}

	if err := ctrl.in.Booking.CancelBooking.CancelBooking(ctx, dto.ID); err != nil {
		return
	}

	ctrl.Response.Data(ctx, ctrl.in.StandardData().NoContent())
}

// EditBooking
// @Summary user edit booking ticket
// @Description user edit booking ticket. ex: change seat
// @Tags 		Booking
// @Accept 		json
// @Produce		json
// @Param		id path string true "id"
// @Param 		params body EditBookingCond true "Request Body"
// @Success 	200 {object} apis.StandardResponse{data=BookingResponse}
// @Failure 	400 {object} apis.StandardResponse{error=apis.StandardError}
// @Failure 	500 {object} apis.StandardResponse{error=apis.StandardError}
// @Router 		/booking/:id [patch]
func (ctrl *Booking) EditBooking(ctx *gin.Context) {
	var err error
	defer func() {
		timelogger.LogTime(ctx)()
		if err != nil {
			ctrl.Response.Data(ctx, ctrl.in.StandardData().BadRequest(nil, err))
		}
	}()

	dto, err := apis.RequestParser[EditBookingCond]().Json().Uri().Bind(ctx)
	if err != nil {
		return
	}

	cond, err := apis.Convert[bo.EditBookingCond](ctx, dto)
	if err != nil {
		return
	}

	if err := ctrl.in.Booking.EditBooking.EditBooking(ctx, cond); err != nil {
		return
	}

	booking, err := ctrl.in.Booking.GetBooking.GetBooking(ctx, cond.ID)
	if err != nil {
		return
	}

	ctrl.Response.Data(ctx, ctrl.in.StandardData().OK(&BookingResponse{
		ID:               booking.ID,
		AirlineCode:      booking.Flight.AirlineCode,
		FlightNumber:     booking.Flight.FlightNumber,
		DepartureAirport: booking.Flight.DepartureAirport,
		ArrivalAirport:   booking.Flight.ArrivalAirport,
		DepartureTime:    booking.Flight.DepartureTime,
		ArrivalTime:      booking.Flight.ArrivalTime,
		Email:            booking.User.Email,
		PhoneCountryCode: booking.User.PhoneCountryCode,
		PhoneNumber:      booking.User.PhoneNumber,
		ClassCode:        booking.Class.ClassCode,
		BaggageAllowance: booking.Class.BaggageAllowance,
		Refundable:       booking.Class.Refundable,
		SeatSelection:    booking.Class.SeatSelection,
		SeatNumber:       lo.Ternary(lo.IsNil(booking.Seat), booking.Seat.SeatNumber, ""),
		Status:           booking.Status,
		Price:            booking.Price,
		ExpiredAt:        booking.ExpiredAt,
	}))
}
