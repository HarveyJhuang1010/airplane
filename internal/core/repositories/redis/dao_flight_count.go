package redis

import (
	"airplane/internal/infrastructure/rediscli"
	"context"
	"errors"
	"fmt"
	"time"
)

type IFlightMaxSellCountDao interface {
	DecrCount(ctx context.Context, id int64, expire time.Duration) (int64, error)
	SetCount(ctx context.Context, id, count int64, expire time.Duration) error
	Exist(ctx context.Context, id int64) (bool, error)
}

func newFlightMaxSellCountDao(in dependence) *FlightMaxSellCountDao {
	return &FlightMaxSellCountDao{
		in:        in,
		keyPrefix: keyPrefixFlight,
	}
}

type FlightMaxSellCountDao struct {
	in        dependence
	keyPrefix string
}

func (dao *FlightMaxSellCountDao) SetCount(ctx context.Context, id, count int64, expire time.Duration) error {
	if _, err := dao.in.RedisClient.Set(ctx, dao.key(id), count, expire).Result(); err != nil {
		return err
	}
	return nil
}

func (dao *FlightMaxSellCountDao) DecrCount(ctx context.Context, id int64, expire time.Duration) (int64, error) {
	var (
		cmd *rediscli.IntCmd
	)

	if _, err := dao.in.RedisClient.TxPipelined(ctx, func(pipe rediscli.Pipeliner) error {
		cmd = pipe.Decr(ctx, dao.key(id))
		pipe.Expire(ctx, dao.key(id), expire)
		return nil
	}); err != nil {
		return 0, dao.in.common.ErrorHandle(err)
	}

	if cmd == nil {
		return 0, errors.New("cmd is nil")
	}
	return cmd.Result()
}

func (dao *FlightMaxSellCountDao) Exist(ctx context.Context, id int64) (bool, error) {
	result, err := dao.in.RedisClient.Exists(ctx, dao.key(id)).Result()
	if err != nil {
		return false, err
	}

	return result == 1, nil
}

func (dao *FlightMaxSellCountDao) key(id int64) string {
	return fmt.Sprintf("%s:%s:%d", dao.keyPrefix, "maxSellCount", id)
}
