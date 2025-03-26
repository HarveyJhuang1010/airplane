package cron

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"airplane/internal/components/logger"
	"airplane/internal/domain/interfaces"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type cronService struct {
	in     digIn
	logger logger.ILogger

	stop chan struct{}
}

func newCronService(in digIn) *cronService {
	return &cronService{
		in:     in,
		logger: in.Logger.SysLogger.Named("cronService"),
		stop:   make(chan struct{}, 1),
	}
}

func (s *cronService) Run(ctx context.Context, stop context.CancelFunc) error {

	defer func() {
		if r := recover(); r != nil {
			s.logger.Error(nil, "panic", zap.Any("panic", r))
		}
		stop()
	}()

	l, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	// initiate cron server and configure with secondOptional for compatibility from previous version
	server := cron.New(
		cron.WithLocation(l),
		cron.WithParser(
			cron.NewParser(
				cron.SecondOptional|
					cron.Minute|
					cron.Hour|
					cron.Dom|
					cron.Month|
					cron.Dow|
					cron.Descriptor,
			),
		),
		cron.WithChain(
			recoverJob(s.logger),
			skipIfStillRunning(s.logger),
		),
	)

	//tasks
	tasks := []interfaces.CronTask{
		// add tasks here
		s.in.CheckBooking,
	}

	for _, task := range tasks {
		s.logger.Info(ctx, "add job", zap.String("schedule", task.Schedule()), zap.String("name", task.Name()))
		if tid, err := server.AddJob(task.Schedule(), task); err != nil {
			s.logger.Error(ctx, err, zap.String("msg", "job added failed"), zap.String("name", task.Name()))
			panic(err)
		} else {
			s.logger.Info(ctx, "job added", zap.String("name", task.Name()), zap.String("tid", fmt.Sprint(tid)))
		}
	}

	server.Start()

	<-s.stop

	return nil
}

func (s *cronService) Shutdown(ctx context.Context) error {
	// TODO: graceful shutdown
	s.stop <- struct{}{}
	s.logger.Info(nil, "shutdown cron")

	return nil
}

// recoverJob panics in wrapped jobs and log them with the provided logger.
func recoverJob(logger logger.ILogger) cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		return cron.FuncJob(func() {
			defer func() {
				if r := recover(); r != nil {
					const size = 64 << 10
					buf := make([]byte, size)
					buf = buf[:runtime.Stack(buf, false)]
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					logger.Error(nil, "panic", zap.ByteString("stack", buf), zap.Error(err))
				}
			}()
			j.Run()
		})
	}
}

// delayIfStillRunning serializes jobs, delaying subsequent runs until the
// previous one is complete. Jobs running after a delay of more than a minute
// have the delay logged at Info.
func delayIfStillRunning(logger *zap.Logger) cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		var mu sync.Mutex
		return cron.FuncJob(func() {
			start := time.Now()
			mu.Lock()
			defer mu.Unlock()
			if dur := time.Since(start); dur > time.Minute {
				logger.Info("delay", zap.Duration("duration", dur))
			}
			j.Run()
		})
	}
}

// skipIfStillRunning skips an invocation of the Job if a previous invocation is
// still running. It logs skips to the given logger at Info level.
func skipIfStillRunning(logger logger.ILogger) cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		var ch = make(chan struct{}, 1)
		ch <- struct{}{}
		return cron.FuncJob(func() {
			select {
			case v := <-ch:
				j.Run()
				ch <- v
			default:
				logger.Debug(nil, "skip")
			}
		})
	}
}
