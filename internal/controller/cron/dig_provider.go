package cron

import (
	"airplane/internal/core/usecase/booking"
	"airplane/internal/domain/interfaces"
	"sync"

	"airplane/internal/components/apis"
	"airplane/internal/components/logger"
	"go.uber.org/dig"
)

var (
	self *packet
)

func NewCronTask(in digIn) digOut {
	self = &packet{}
	self.Do(func() {
		self.in = in
		self.digOut = digOut{
			CheckBooking: newCheckBookingTask(in),
		}
	})

	return self.digOut
}

type digIn struct {
	dig.In

	Logger   *logger.Loggers
	Response apis.IResponse

	// core
	Booking *booking.Usecase
}

type packet struct {
	sync.Once

	in digIn

	digOut
}

type digOut struct {
	dig.Out

	CheckBooking interfaces.CronTask `name:"check_booking"`
}
