package mqcli

import (
	"airplane/internal/components/logger"
	"airplane/internal/config"
	"sync"

	"go.uber.org/dig"
)

var (
	self *set
)

func NewMQClient(in digIn) digOut {
	self = &set{}
	self.Do(func() {
		self.in = in
		self.digOut = digOut{
			Kafka: newKafka(in),
		}
	})

	return self.digOut
}

type digIn struct {
	dig.In

	Config *config.Config
	Logger *logger.Loggers
}

type set struct {
	sync.Once
	in digIn

	digOut
}

type digOut struct {
	dig.Out

	Kafka *Kafka
}
