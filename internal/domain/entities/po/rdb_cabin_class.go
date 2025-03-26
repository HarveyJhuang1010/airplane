package po

import (
	"airplane/internal/enum"
	"github.com/shopspring/decimal"
)

type CabinClass struct {
	ID               int64               `gorm:"primaryKey" json:"id"`
	FlightID         int64               `gorm:"not null;index" json:"flightID"`
	ClassCode        enum.CabinClassCode `gorm:"type:varchar(20);not null" json:"classCode"`
	Price            decimal.Decimal     `gorm:"type:decimal(10,2);not null" json:"price"`
	BaggageAllowance int                 `gorm:"not null;default:20" json:"baggageAllowance"`
	Refundable       bool                `gorm:"default:false" json:"refundable"`
	SeatSelection    bool                `gorm:"default:false" json:"seatSelection"`
	MaxSeats         int                 `gorm:"not null" json:"maxSeats"`
	RemainSeats      int                 `gorm:"not null" json:"remainSeats"`

	Flight Flight  `gorm:"foreignKey:FlightID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Seats  []*Seat `gorm:"foreignKey:CabinClassID" json:"-"`

	At
}

func (CabinClass) TableName() string {
	return "cabin_class"
}
