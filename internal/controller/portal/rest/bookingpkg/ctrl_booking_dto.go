package bookingpkg

import (
	"airplane/internal/enum"
	"github.com/shopspring/decimal"
	"time"
)

type BookingResponse struct {
	ID               int64               `json:"id,string"`
	AirlineCode      string              `json:"airlineCode"`
	FlightNumber     string              `json:"flightNumber"`
	DepartureAirport string              `json:"departureAirport"`
	ArrivalAirport   string              `json:"arrivalAirport"`
	DepartureTime    time.Time           `json:"departureTime"`
	ArrivalTime      time.Time           `json:"arrivalTime"`
	Email            string              `json:"email"`
	PhoneCountryCode string              `json:"phoneCountryCode"`
	PhoneNumber      string              `json:"phoneNumber"`
	ClassCode        enum.CabinClassCode `json:"classCode" enums:"economy_standard,economy_flex,business_basic,business_standard" swaggertype:"string"`
	BaggageAllowance int                 `json:"baggageAllowance,string"`
	Refundable       bool                `json:"refundable"`
	SeatSelection    bool                `json:"seatSelection"`
	SeatNumber       string              `json:"seatNumber"`
	Status           enum.BookingStatus  `json:"status" enums:"pending,confirming,confirmed,failed,cancelled" swaggertype:"string"`
	Price            decimal.Decimal     `json:"price"`
	ExpiredAt        time.Time           `json:"expiredAt" time_format:"2006-01-02T15:04:05Z07:00"`
}

type AddBookingResponse struct {
	ID int64 `json:"id"`
}

type AddBookingCond struct {
	ID           int64  `json:"id"`
	FlightID     int64  `json:"flightID"`
	CabinClassID int64  `json:"cabinClassID"`
	SeatID       *int64 `json:"seatID"`
	Email        string `json:"email"`
	CountryCode  string `json:"countryCode"`
	PhoneNumber  string `json:"phoneNumber"`
}

type CancelBookingCond struct {
	ID int64 `uri:"id" swaggerignore:"true"`
}

type EditBookingCond struct {
	ID           int64 `uri:"id" swaggerignore:"true"`
	CabinClassID int64 `json:"cabinClassID"`
	SeatID       int64 `json:"seatID"`
}

type GetBookingCond struct {
	ID int64 `uri:"id" swaggerignore:"true"`
}
