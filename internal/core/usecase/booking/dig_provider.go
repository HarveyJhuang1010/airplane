package booking

import (
	"airplane/internal/components/logger"
	"airplane/internal/core/business/payment"
	"airplane/internal/core/business/seat"
	"airplane/internal/core/business/user"
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/core/repositories/redis"
	"airplane/internal/infrastructure/mqcli"
	"airplane/internal/infrastructure/snowflake"
	"go.uber.org/dig"
	"sync"
)

var self *packet

func New(in digIn) digOut {
	self = &packet{}
	self.Do(func() {
		self.in = in
		self.digOut = digOut{
			Usecase: newBookingUsecase(in),
		}
	})

	return self.digOut
}

type packet struct {
	sync.Once

	in digIn

	digOut
}

type digIn struct {
	dig.In

	Logger          *logger.Loggers
	Snowflake       *snowflake.Snowflake
	DBRepository    *rdb.Repository
	RedisRepository *redis.Repository
	Kafka           *mqcli.Kafka
	Payment         *payment.Usecase
	User            *user.Usecase
	Seat            *seat.Usecase
}

type digOut struct {
	dig.Out

	Usecase *Usecase
}

func newBookingUsecase(in digIn) *Usecase {
	return &Usecase{
		AddBooking:        newAddBooking(in),
		CancelBooking:     newCancelBooking(in),
		CheckBooking:      newCheckBooking(in),
		ConfirmBooking:    newConfirmBooking(in),
		GetBooking:        newGetBooking(in),
		HandleOverBooking: newHandleOverBooking(in),
		EditBooking:       newEditBooking(in),
	}
}

type Usecase struct {
	AddBooking        *AddBooking
	CancelBooking     *CancelBooking
	CheckBooking      *CheckBooking
	ConfirmBooking    *ConfirmBooking
	GetBooking        *GetBooking
	HandleOverBooking *HandleOverBooking
	EditBooking       *EditBooking
}
