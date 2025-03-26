package listener

import (
	"airplane/internal/components/logger"
	"airplane/internal/config"
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/core/repositories/redis"
	"airplane/internal/core/usecase/booking"
	"airplane/internal/domain/interfaces"
	"airplane/internal/infrastructure/mqcli"
	"airplane/internal/infrastructure/snowflake"
	"sync"

	"go.uber.org/dig"
)

var (
	self *packet
)

func NewListener(in digIn) digOut {
	self = &packet{}
	self.Do(func() {
		self.in = in
		self.digOut = digOut{
			AddBookingListener:     newAddBookingListener(in),
			ConfirmBookingListener: newConfirmBookingListener(in),
		}
	})

	return self.digOut
}

type digIn struct {
	dig.In

	// 套件
	AppConf   *config.Config
	Logger    *logger.Loggers
	Snowflake *snowflake.Snowflake

	// DB
	DBRepository    *rdb.Repository
	RedisRepository *redis.Repository

	// 其他
	Kafka   *mqcli.Kafka
	Booking *booking.Usecase
}

type packet struct {
	sync.Once

	in digIn

	digOut
}

type digOut struct {
	dig.Out

	AddBookingListener     interfaces.Listener `name:"add_booking_listener"`
	ConfirmBookingListener interfaces.Listener `name:"confirm_booking_listener"`
}
