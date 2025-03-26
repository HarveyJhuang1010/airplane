package po

import (
	"airplane/internal/enum"
	"time"
)

type Flight struct {
	ID               int64             `gorm:"primaryKey" json:"id"`
	AirlineCode      string            `gorm:"type:varchar(10);not null" json:"airlineCode"`
	FlightNumber     string            `gorm:"type:varchar(20);not null" json:"flightNumber"`
	DepartureAirport string            `gorm:"type:varchar(10);not null;index" json:"departureAirport"`
	ArrivalAirport   string            `gorm:"type:varchar(10);not null;index" json:"arrivalAirport"`
	DepartureTime    time.Time         `gorm:"not null;index" json:"departureTime"`
	ArrivalTime      time.Time         `gorm:"not null" json:"arrivalTime"`
	TotalSeats       int               `gorm:"not null" json:"totalSeats"`
	OverbookingLimit int               `gorm:"default:5" json:"overbookingLimit"`
	SellableSeats    int               `gorm:"not null" json:"sellableSeats"`
	Status           enum.FlightStatus `gorm:"type:varchar(20);not null" json:"status"`

	CabinClasses []*CabinClass `gorm:"foreignKey:FlightID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`

	At
}

func (Flight) TableName() string {
	return "flight"
}

type FlightListCond struct {
	*Pager
	DepartureAirport     *string
	ArrivalAirport       *string
	DepartureTimeStartAt *time.Time
	DepartureTimeEndAt   *time.Time
	PreloadCabinClasses  bool
	Status               []enum.FlightStatus
	CanSell              bool
}
