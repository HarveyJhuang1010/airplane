package rediscli

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"airplane/internal/errs"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type RedisCmd interface {
	*redis.StatusCmd |
		*redis.DurationCmd |
		*redis.SliceCmd |
		*redis.FloatCmd |
		*redis.ScanCmd |
		*redis.MapStringStringCmd |
		*redis.MapStringIntCmd |
		*redis.KeyValueSliceCmd |
		*redis.CommandsInfoCmd |
		*redis.BoolCmd |
		*redis.BoolSliceCmd |
		*redis.IntCmd |
		*redis.IntSliceCmd |
		*redis.StringCmd |
		*redis.StringSliceCmd
}

// Wrap from given error return mapping value
func Wrap[T RedisCmd](val T) T {
	v := any(val).(redis.Cmder)
	if v.Err() != nil {
		v.SetErr(WrapError(v.Err()))
	}

	return v.(T)
}

func WrapError(err error) error {
	if errors.Is(err, redis.Nil) {
		return errs.ErrRedisNotFound
	} else if errors.Is(err, redis.TxFailedErr) {
		return errs.ErrRedisTxFailed
	} else {
		return fmt.Errorf("%w, %s", errs.ErrRedisOperationFailed, err)
	}
}

func GetDataByKey(ctx context.Context, r *Redis, key string, result any) (bool, error) {
	exist, err := Wrap(r.Exists(ctx, key)).Result()
	if err != nil {
		return false, errors.Wrapf(errs.ErrRedisOperationFailed, "getDataByKey ExistsByAccount failed: %v", err)
	}

	if exist == 0 {
		return false, nil
	}

	data, err := Wrap(r.Get(ctx, key)).Bytes()
	if err != nil {
		return false, errors.Wrapf(errs.ErrRedisOperationFailed, "getDataByKey GetByAccount failed: %v", err)
	}

	if err := json.Unmarshal(data, result); err != nil {
		return false, errors.Wrapf(errs.ErrRedisOperationFailed, "getDataByKey Unmarshal failed: %v", err)
	}

	return true, nil
}

func SetDataByKey(ctx context.Context, r *Redis, key string, result any, duration time.Duration) error {
	if err := Wrap(r.Set(ctx, key, result, duration)).Err(); err != nil {
		return errors.Wrapf(errs.ErrRedisOperationFailed, "setDataByKey GetByAccount failed: %v", err)
	}

	return nil
}
