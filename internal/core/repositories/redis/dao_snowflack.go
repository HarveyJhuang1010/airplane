package redis

import (
	"airplane/internal/infrastructure/redlock"
	"context"
	"strconv"
	"time"
)

type ISnowflakeDao interface {
	Lock(ctx context.Context, idx int) (Unlock func() error, err error)
}

func newSnowflakeDao(in dependence) *snowflakeDao {
	return &snowflakeDao{
		in:        in,
		keyPrefix: keyPrefixSnowflake,
	}
}

type snowflakeDao struct {
	in        dependence
	keyPrefix string
}

// Lock Redis Lock
func (dao *snowflakeDao) Lock(ctx context.Context, idx int) (Unlock func() error, err error) {
	redLock := dao.in.RedLock.NewMutex(ctx, dao.lockKey(idx), redlock.WithExpiry(5*time.Second))
	if err := redLock.Lock(); err != nil {
		return nil, err
	}
	return redLock.Unlock, nil
}

func (dao *snowflakeDao) lockKey(idx int) string {
	return dao.keyPrefix + ":lock:" + strconv.Itoa(idx)
}
