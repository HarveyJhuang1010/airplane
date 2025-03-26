package bo

import (
	"airplane/internal/enum"
	"github.com/shopspring/decimal"
	"time"
)

type FlightDetail struct {
	Flight
	SellableSeats int           `json:"sellableSeats"`
	CabinClasses  []*CabinClass `json:"cabinClasses"`
}

type ListFlightCond struct {
	*Pager
	DepartureAirport     *string
	ArrivalAirport       *string
	DepartureTimeStartAt *time.Time
	DepartureTimeEndAt   *time.Time
}

type Flight struct {
	ID               int64             `json:"id"`
	AirlineCode      string            `json:"airlineCode"`
	FlightNumber     string            `json:"flightNumber"`
	DepartureAirport string            `json:"departureAirport"`
	ArrivalAirport   string            `json:"arrivalAirport"`
	DepartureTime    time.Time         `json:"departureTime"`
	ArrivalTime      time.Time         `json:"arrivalTime"`
	Status           enum.FlightStatus `json:"status"`
}

type CabinClass struct {
	ID               int64
	ClassCode        enum.CabinClassCode
	Price            decimal.Decimal
	BaggageAllowance int
	Refundable       bool
	SeatSelection    bool
	MaxSeats         int
	RemainSeats      int
}
