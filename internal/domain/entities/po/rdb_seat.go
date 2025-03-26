package po

import (
	"airplane/internal/enum"
)

type Seat struct {
	ID           int64           `gorm:"primaryKey" json:"id"`
	FlightID     int64           `gorm:"not null;index" json:"flightID"`
	CabinClassID int64           `gorm:"index" json:"cabinClassID"`
	SeatNumber   string          `gorm:"type:varchar(5);not null" json:"seatNumber"`
	Status       enum.SeatStatus `gorm:"type:varchar(20);not null" json:"status"`

	CabinClass CabinClass `gorm:"foreignKey:CabinClassID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Flight     Flight     `gorm:"foreignKey:FlightID"`

	At
}

func (Seat) TableName() string {
	return "seat"
}

type SeatListCond struct {
	*Pager
	FlightID *int64
	Status   *enum.SeatStatus
}

type SeatUpdateCond struct {
	ID     int64
	Status enum.SeatStatus
}
