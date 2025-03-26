package redis

import (
	"airplane/internal/errs"
	"errors"
	"github.com/redis/go-redis/v9"
)

type common struct {
}

func (common) ErrorHandle(err error) error {
	switch {
	case errors.Is(err, redis.Nil):
		return errs.ErrRecordNotFound.Trace()
	}

	return err
}
