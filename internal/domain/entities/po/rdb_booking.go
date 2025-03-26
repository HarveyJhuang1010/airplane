package po

import (
	"airplane/internal/enum"
	"github.com/shopspring/decimal"
	"time"
)

type Booking struct {
	ID           int64              `gorm:"primaryKey" json:"id"`
	FlightID     int64              `gorm:"not null;index" json:"flightID"`
	UserID       int64              `gorm:"not null;index" json:"userID"`
	CabinClassID int64              `gorm:"not null" json:"cabinClassID"`
	SeatID       *int64             `gorm:"index" json:"seatID"`
	Status       enum.BookingStatus `gorm:"type:varchar(20);not null" json:"status"`
	Price        decimal.Decimal    `gorm:"type:decimal(10,2);not null" json:"price"`
	ExpiredAt    time.Time          `gorm:"not null" json:"expiredAt"`

	Flight Flight     `gorm:"foreignKey:FlightID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	User   User       `gorm:"foreignKey:UserID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Class  CabinClass `gorm:"foreignKey:CabinClassID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Seat   *Seat      `gorm:"foreignKey:SeatID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL"`

	At
}

func (Booking) TableName() string {
	return "booking"
}

type BookingListCond struct {
	*Pager
	FlightID  *int64
	UserID    *int64
	PaymentID *int64
	Status    []enum.BookingStatus
}

type BookingUpdateCond struct {
	ID     int64
	Status enum.BookingStatus
}

type BookingUpdateSeatCond struct {
	ID           int64
	SeatID       *int64
	CabinClassID *int64
	Price        *decimal.Decimal
}
