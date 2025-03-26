package redis

import (
	"airplane/internal/infrastructure/rediscli"
	"airplane/internal/infrastructure/redlock"
	"go.uber.org/dig"
)

const (
	keyPrefixSnowflake = "repo:snowflake"
	keyPrefixFlight    = "repo:flight"
)

func New(in digIn) digOut {
	dep := dependence{
		digIn: in,
	}
	dep.RedisClient = redisClient{Redis: in.RedisClient}

	self := &digProvider{
		in: dep,
		out: digOut{
			Repository:   newRepository(dep),
			SnowflakeDao: newSnowflakeDao(dep),
		},
	}

	return self.out
}

type digProvider struct {
	in  dependence
	out digOut
}

type dependence struct {
	digIn

	RedisClient redisClient
	common      common
}

type digIn struct {
	dig.In

	RedLock     redlock.IRedLock
	RedisClient *rediscli.Redis
}

type digOut struct {
	dig.Out

	Repository   *Repository
	SnowflakeDao ISnowflakeDao
}

func newRepository(in dependence) *Repository {
	return &Repository{
		FlightMaxSellCount: newFlightMaxSellCountDao(in),
	}
}

type Repository struct {
	FlightMaxSellCount IFlightMaxSellCountDao
}
