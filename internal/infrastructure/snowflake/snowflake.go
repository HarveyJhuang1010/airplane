package snowflake

import (
	"airplane/internal/components/logger"
	"airplane/internal/tools/rand"
	"context"
	"github.com/bwmarrin/snowflake"
	"time"
)

func newSnowflake(in digIn) *Snowflake {
	return &Snowflake{
		in:     in,
		logger: in.Logger.AppLogger.Named("snowflake"),
	}
}

type Snowflake struct {
	in     digIn
	logger logger.ILogger
	unlock func() error

	*snowflake.Node
}

func (s *Snowflake) Run(ctx context.Context, stop context.CancelFunc) error {
	s.initialize(ctx)
	s.logger.Info(ctx, "Snowflake is running")
	return nil
}

func (s *Snowflake) Shutdown(ctx context.Context) error {
	if s.unlock != nil {
		s.unlock()
	}
	s.logger.Info(ctx, "Snowflake is shutdown")
	return nil
}

func (s *Snowflake) initialize(ctx context.Context) *Snowflake {
	var (
		nodeIdx int
	)

	ctxWithTimeout, _ := context.WithTimeout(ctx, 5*time.Second)

	for {
		select {
		case <-ctxWithTimeout.Done():
			s.logger.Panic(ctx, "initialize snowflake timeout")
		default:
		}

		// 範圍 0 ~ 1023
		nodeIdx = rand.Intn(255)
		un, err := s.in.SnowflakeDao.Lock(ctx, nodeIdx)
		if err != nil {
			continue
		}
		s.unlock = un
		break
	}

	// 起始時間(Epoch)改為 2023-01-01 00:00:00 UTC
	// 上線後請勿調整，定啖好會重複
	// 理順上雪花可以使用 69 年，到 2092 年
	snowflake.Epoch = 1672502400000
	node, err := snowflake.NewNode(int64(nodeIdx))
	if err != nil {
		s.logger.Panic(ctx, err)
	}

	s.Node = node

	return s
}
