package qworker

import (
	"airplane/internal/components/logger"
	"airplane/internal/domain/interfaces"
	"context"
	"go.uber.org/zap"
)

type qworkerService struct {
	in     digIn
	logger logger.ILogger

	stop chan struct{}
}

func newCronService(in digIn) *qworkerService {
	return &qworkerService{
		in:     in,
		logger: in.Logger.SysLogger.Named("qworkerService"),
		stop:   make(chan struct{}, 1),
	}
}

func (s *qworkerService) Run(ctx context.Context, stop context.CancelFunc) error {

	defer func() {
		if r := recover(); r != nil {
			s.logger.Error(nil, "panic", zap.Any("panic", r))
		}
		stop()
	}()

	//tasks
	tasks := []interfaces.Listener{
		// add tasks here
		s.in.AddBookingListener,
		s.in.ConfirmBookingListener,
	}

	for _, task := range tasks {
		s.logger.Info(ctx, "add listener", zap.String("name", task.Name()))
		go task.Listen(ctx)
	}

	<-s.stop

	return nil
}

func (s *qworkerService) Shutdown(ctx context.Context) error {
	// TODO: graceful shutdown
	s.stop <- struct{}{}
	s.logger.Info(nil, "shutdown qworker")

	return nil
}
