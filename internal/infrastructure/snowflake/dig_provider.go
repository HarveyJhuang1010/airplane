package snowflake

import (
	"airplane/internal/components/logger"
	"airplane/internal/config"
	"airplane/internal/core/repositories/redis"
	"sync"

	"go.uber.org/dig"
)

var (
	self *set
)

func New(in digIn) digOut {
	self = &set{}
	self.Do(func() {
		self.in = in
		self.digOut = digOut{
			Snowflake: newSnowflake(in),
		}
	})

	return self.digOut
}

type digIn struct {
	dig.In

	Config       *config.Config
	Logger       *logger.Loggers
	SnowflakeDao redis.ISnowflakeDao
}

type set struct {
	sync.Once
	in digIn

	digOut
}

type digOut struct {
	dig.Out

	Snowflake *Snowflake
}
