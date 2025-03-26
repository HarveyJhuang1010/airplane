package bo

import (
	"airplane/internal/enum"
	"github.com/shopspring/decimal"
	"time"
)

type Booking struct {
	ID           int64              `json:"id"`
	FlightID     int64              `json:"flightID"`
	UserID       int64              `json:"userID"`
	CabinClassID int64              `json:"cabinClassID"`
	SeatID       *int64             `json:"seatID"`
	Status       enum.BookingStatus `json:"status"`
	Price        decimal.Decimal    `json:"price"`
	ExpiredAt    time.Time          `json:"expiredAt"`
	Flight       Flight
	User         User
	Class        CabinClass
	Seat         *Seat
}

type AddBookingCond struct {
	ID           int64
	FlightID     int64
	CabinClassID int64
	SeatID       *int64
	Email        string
	CountryCode  string
	PhoneNumber  string
}

type ConfirmBookingCond struct {
	ID int64
}

type EditBookingCond struct {
	ID           int64
	CabinClassID int64 `json:"cabinClassID"`
	SeatID       int64 `json:"seatID"`
}
