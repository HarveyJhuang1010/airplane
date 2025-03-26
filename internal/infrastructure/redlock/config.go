package redlock

import (
	"time"
)

type config struct {
	expiry time.Duration

	tries int
	delay time.Duration
}

type Option interface {
	Apply(*config)
}

type OptionFunc func(*config)

func (f OptionFunc) Apply(conf *config) {
	f(conf)
}

func WithExpiry(expiry time.Duration) Option {
	return OptionFunc(func(m *config) {
		m.expiry = expiry
	})
}

func WithDelay(tries int) Option {
	return OptionFunc(func(m *config) {
		m.tries = tries
	})
}

func WithTries(delay time.Duration) Option {
	return OptionFunc(func(m *config) {
		m.delay = delay
	})
}
