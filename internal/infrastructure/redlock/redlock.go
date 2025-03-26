package redlock

import (
	"context"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"runtime"
	"strings"
	"time"
)

func newRedLock(param dependence) *redLock {
	return &redLock{param: param}
}

type IRedLock interface {
	NewMutex(ctx context.Context, key string, opts ...Option) *Mutex
}

type redLock struct {
	param dependence
}

// NewMutex 包裝 redsync.NewMutex，實作 watchDog
func (r *redLock) NewMutex(ctx context.Context, key string, opts ...Option) *Mutex {
	config := &config{
		expiry: 1 * time.Second,        // 預設 1 秒過期
		tries:  32,                     // 預設嘗試 32 次
		delay:  200 * time.Millisecond, // 預設延遲 200 毫秒
	}
	for _, opt := range opts {
		opt.Apply(config)
	}

	ctx, cancel := context.WithCancel(ctx)

	return &Mutex{
		ctx:    ctx,
		cancel: cancel,
		mx: redsync.New(goredis.NewPool(r.param.in.RedisClient.Client)).
			NewMutex(
				key,
				redsync.WithExpiry(config.expiry), // 鎖的過期時間
				redsync.WithTries(config.tries),
				redsync.WithRetryDelay(config.delay),
			),
		config: config,
	}
}

type Mutex struct {
	ctx    context.Context
	cancel context.CancelFunc

	mx     *redsync.Mutex
	config *config
}

func (m *Mutex) Lock() error {
	if err := m.mx.LockContext(m.ctx); err != nil {
		return err
	}

	go m.watchDog()

	return nil
}

func (m *Mutex) UntilLock() error {
	for {
		select {
		case <-m.ctx.Done():
			return nil
		default:
		}
		if err := m.Lock(); err != nil {
			if strings.Contains(err.Error(), "lock already taken") {
				runtime.Gosched()
				continue
			}
			return err
		}
		break
	}

	return nil
}

func (m *Mutex) Unlock() error {
	defer m.cancel()
	_, err := m.mx.UnlockContext(m.ctx)
	return err
}

func (m *Mutex) Extend() (bool, error) {
	return m.mx.ExtendContext(m.ctx)
}

func (m *Mutex) watchDog() {
	t := time.NewTicker(m.config.expiry / 2)
	for range t.C {
		select {
		case <-m.ctx.Done():
			return
		default:
		}

		if _, err := m.mx.ExtendContext(m.ctx); err != nil {
			m.cancel()
			return
		}
	}
}
