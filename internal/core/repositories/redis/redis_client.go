package redis

import (
	"airplane/internal/infrastructure/rediscli"
	"context"
	"github.com/golang/groupcache/singleflight"
	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	*rediscli.Redis
	single singleflight.Group
}

func (c *redisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	data, _ := c.single.Do(key, func() (interface{}, error) {
		data := c.Redis.Get(ctx, key)
		return data, nil
	})

	if d, ok := data.(*redis.StringCmd); ok {
		return d
	}
	return nil
}
