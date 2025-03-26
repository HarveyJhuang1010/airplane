package bo

import "airplane/internal/enum"

type Seat struct {
	ID           int64           `json:"id"`
	FlightID     int64           `json:"flightID"`
	CabinClassID int64           `json:"cabinClassID"`
	SeatNumber   string          `json:"seatNumber"`
	Status       enum.SeatStatus `json:"status"`
}
