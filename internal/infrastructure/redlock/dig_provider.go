package redlock

import (
	"airplane/internal/components/logger"
	"airplane/internal/infrastructure/rediscli"
	"go.uber.org/dig"
)

func New(in digIn) IRedLock {
	dependence := dependence{in: in}
	return newRedLock(dependence)
}

type digIn struct {
	dig.In

	RedisClient *rediscli.Redis
	Logger      *logger.Loggers
}

type dependence struct {
	in digIn
}
