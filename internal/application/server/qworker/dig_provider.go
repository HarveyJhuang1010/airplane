package qworker

import (
	"sync"

	"airplane/internal/components/apis"
	"airplane/internal/components/launcher"
	"airplane/internal/components/logger"
	"airplane/internal/config"
	"airplane/internal/domain/interfaces"
	"go.uber.org/dig"
)

var (
	self *packet
)

func NewServer(in digIn) digOut {
	self = &packet{}
	self.Do(func() {
		self.in = in
		self.digOut = digOut{
			ApiService:  newApiService(in),
			CronService: newCronService(in),
		}
	})

	return self.digOut
}

type digIn struct {
	dig.In

	Logger *logger.Loggers
	Config *config.Config
	Apis   *apis.Apis

	// tasks
	AddBookingListener     interfaces.Listener `name:"add_booking_listener"`
	ConfirmBookingListener interfaces.Listener `name:"confirm_booking_listener"`
}

type packet struct {
	sync.Once

	in digIn

	digOut
}

type digOut struct {
	dig.Out

	ApiService  *apis.Service     `name:"qworker_srv_api_service"`
	CronService launcher.IService `name:"qworker_srv_qworker_service"`
}
