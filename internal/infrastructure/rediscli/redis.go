package rediscli

import (
	"airplane/internal/components/logger"
	"airplane/internal/config"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

//const Nil = redis.Nil

type Pipeliner = redis.Pipeliner
type Cmdable = redis.Cmdable
type IntCmd = redis.IntCmd

// 暫時保留 start
//var (
//	defaultClient Redis
//)
//
//func GetRedis() *Redis {
//	return &defaultClient
//}

// 暫時保留 end

func newRedis(in digIn) *Redis {
	return &Redis{
		config: in.Config.Redis,
		logger: in.Logger.SysLogger.Named("redis"),
	}
}

type Redis struct {
	config *config.RedisConfig
	*redis.Client

	logger logger.ILogger
}

func (r *Redis) GetConfig() *config.RedisConfig {
	return r.config
}

func (r *Redis) Run(ctx context.Context, stop context.CancelFunc) error {
	r.initialize(ctx, r.config)
	r.logger.Info(ctx, "Redis is running")
	return nil
}

func (r *Redis) Shutdown(ctx context.Context) error {
	r.Close()
	r.logger.Info(ctx, "Redis is shutdown")
	return nil
}

func (r *Redis) initialize(ctx context.Context, cfg *config.RedisConfig) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	r.logger.Debug(ctx, "Initializing redis ...")
	defer r.logger.Debug(ctx, "Done")

	r.config = cfg
	r.Client = client
}

// InitTestRedis init test redis
func (r *Redis) InitTestRedis(ctx context.Context, cfg *config.RedisConfig) (*Redis, error) {
	cfg.DB = 15
	r.initialize(ctx, cfg)
	if r.Client == nil {
		return nil, errors.New("init redis client failed")
	}

	return r, nil
}

func (r *Redis) WatchKeys(ctx context.Context, fn func(context.Context) error, keys []string, retries ...int) (err error) {

	maxRetries := 10
	if len(retries) > 0 {
		maxRetries = retries[0]
	}
	txFunc := func(tx *redis.Tx) error {
		return fn(injectTx(ctx, tx))
	}

	for i := 0; i < maxRetries; i++ {
		err := r.Watch(ctx, txFunc, keys...)
		if err == nil {
			// Success.
			return nil
		}
		if err == redis.TxFailedErr {
			// Optimistic lock lost. Retry.
			continue
		}
		// Return any other error.
		return err
	}
	return nil
}

type txKey struct{}

func (r *Redis) Tx(ctx context.Context) *redis.Tx {
	tx := extractTx(ctx)
	if tx != nil {
		return tx
	}
	return nil
}

// injectTx injects transaction to context
func injectTx(ctx context.Context, tx *redis.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// extractTx extracts transaction from context
func extractTx(ctx context.Context) *redis.Tx {
	if tx, ok := ctx.Value(txKey{}).(*redis.Tx); ok {
		return tx
	}
	return nil
}
