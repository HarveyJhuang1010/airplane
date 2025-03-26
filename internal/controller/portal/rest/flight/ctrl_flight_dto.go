package flight

import (
	"airplane/internal/components/apis"
	"airplane/internal/enum"
	"github.com/shopspring/decimal"
	"time"
)

type FlightResponse struct {
	// Flight ID
	ID int64 `json:"id,string"`
	// 航空公司代碼
	AirlineCode string `json:"airlineCode"`
	// 航班號碼
	FlightNumber string `json:"flightNumber"`
	// 出發機場
	DepartureAirport string `json:"departureAirport"`
	// 到達機場
	ArrivalAirport string `json:"arrivalAirport"`
	// 出發時間
	DepartureTime time.Time `json:"departureTime" time_format:"2006-01-02T15:04:05Z07:00"`
	// 到達時間
	ArrivalTime time.Time `json:"arrivalTime" time_format:"2006-01-02T15:04:05Z07:00"`
	// 航班狀態
	Status enum.FlightStatus `json:"status" enums:"scheduled,boarding,departed,arrived,cancelled" swaggertype:"string"`
	// 可售座位數
	SellableSeats int `json:"sellableSeats,string"`
	// 艙等
	CabinClasses []*CabinClassResponse `json:"cabinClasses"`
}

type FlightListRequest struct {
	*apis.Pager

	// 出發機場
	DepartureAirport *string `form:"departureAirport" swaggertype:"string"`
	// 到達機場
	ArrivalAirport *string `form:"arrivalAirport" swaggertype:"string"`
	// 出發時間起始(YYYY-MM-DDTHH:MM:SSZ)
	DepartureTimeStartAt *time.Time `form:"departureTimeStartAt" swaggertype:"string"  time_format:"2006-01-02T15:04:05Z07:00"`
	// 出發時間結束(YYYY-MM-DDTHH:MM:SSZ)
	DepartureTimeEndAt *time.Time `form:"departureTimeEndAt" swaggertype:"string"  time_format:"2006-01-02T15:04:05Z07:00"`
}

type CabinClassResponse struct {
	ID               int64               `json:"id,string"`
	ClassCode        enum.CabinClassCode `json:"classCode" enums:"economy_standard,economy_flex,business_basic,business_standard" swaggertype:"string"`
	Price            decimal.Decimal     `json:"price"`
	BaggageAllowance int                 `json:"baggageAllowance,string"`
	Refundable       bool                `json:"refundable"`
	SeatSelection    bool                `json:"seatSelection"`
	MaxSeats         int                 `json:"maxSeats,string"`
	RemainSeats      int                 `json:"remainSeats,string"`
}
