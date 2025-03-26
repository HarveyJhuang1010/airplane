package cron

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
	CheckBooking interfaces.CronTask `name:"check_booking"`
}

type packet struct {
	sync.Once

	in digIn

	digOut
}

type digOut struct {
	dig.Out

	ApiService  *apis.Service     `name:"cron_srv_api_service"`
	CronService launcher.IService `name:"cron_srv_cron_service"`
}
