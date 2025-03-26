package launcher

import (
	"airplane/internal/components/logger"
	"airplane/internal/config"
	"context"
	"fmt"
	"go.uber.org/dig"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"time"
)

type digIn struct {
	dig.In

	Config *config.Config
	Logger *logger.Loggers
}

func NewLauncher(in digIn) *Launcher {
	return &Launcher{
		in:              in,
		serverInterrupt: newServerInterrupt(in),
	}
}

type Launcher struct {
	in digIn

	once            sync.Once
	serverInterrupt *serverInterrupt
}

type service struct {
	IService
	async bool
}

// Infrastructure 用於啟動服務後，啟動基礎設施服務。在 shutdown 時，會最後關閉基礎設施服務
func (s *Launcher) Infrastructure(ctx context.Context, stop context.CancelFunc, services ...IService) {
	serves := make([]*service, len(services))
	for i, srv := range services {
		serves[i] = &service{IService: srv, async: false}
	}

	s.serverInterrupt.RegisterInfrastructure(serves...)

	s.launcher(ctx, stop, serves...)
}

// ListenAndServe 用於啟動服務後，啟動服務
func (s *Launcher) ListenAndServe(ctx context.Context, stop context.CancelFunc, services ...IService) {
	serves := make([]*service, len(services))
	for i, srv := range services {
		serves[i] = &service{IService: srv, async: true}
	}
	s.serverInterrupt.RegisterService(serves...)

	s.launcher(ctx, stop, serves...)
}

func (s *Launcher) launcher(ctx context.Context, stop context.CancelFunc, services ...*service) {
	s.serverInterrupt.Listen(ctx)

	for _, serv := range services {
		if serv.async {
			go s.run(serv.IService, ctx, stop)
		} else {
			s.run(serv.IService, ctx, stop)
		}
		s.in.Logger.SysLogger.Info(ctx, fmt.Sprintf("start %s", reflect.TypeOf(serv.IService).String()))
	}

}

func (s *Launcher) run(service IService, ctx context.Context, stop context.CancelFunc) {
	if err := service.Run(ctx, stop); err != nil {
		fmt.Printf("run error: %v \n", err)
		s.in.Logger.SysLogger.Panic(ctx, err)
		os.Exit(1)
	}
}

type IService interface {
	Run(ctx context.Context, stop context.CancelFunc) error
	Shutdown(ctx context.Context) error
}

func newServerInterrupt(in digIn) *serverInterrupt {
	return &serverInterrupt{
		in: in,
	}
}

type serverInterrupt struct {
	in   digIn
	mx   sync.Mutex
	once sync.Once

	infrastructures []*service
	services        []*service
}

func (s *serverInterrupt) RegisterInfrastructure(services ...*service) *serverInterrupt {
	s.mx.Lock()
	s.infrastructures = append(s.services, services...)
	s.mx.Unlock()

	return s
}

func (s *serverInterrupt) RegisterService(services ...*service) *serverInterrupt {
	s.mx.Lock()
	s.services = append(s.services, services...)
	s.mx.Unlock()

	return s
}

func (s *serverInterrupt) Listen(ctx context.Context) {
	s.once.Do(func() {
		go s.listen(ctx)
	})
}

func (s *serverInterrupt) listen(ctx context.Context) {
	exception := make(chan bool, 1)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	go func(ctx context.Context) {
		c := <-interrupt

		s.mx.Lock()
		defer s.mx.Unlock()
		s.in.Logger.SysLogger.Info(ctx, fmt.Sprintf("Server Shutdown, osSignal: %v", c))
		s.in.Logger.SysLogger.Info(ctx, fmt.Sprintf("Shutdown Server ..."))

		// TODO: shutdown 超時之後設定檔調整
		ctxWithTimeout, _ := context.WithTimeout(context.Background(), 30*time.Second)
		go func(ctx context.Context) {
			select {
			case <-ctx.Done():
				s.in.Logger.SysLogger.Error(ctx, "Server exiting timeout")
				exception <- true
				return
			}
		}(ctxWithTimeout)

		var isException bool
		isException = isException || s.shutdown(ctxWithTimeout, s.services)
		isException = isException || s.shutdown(ctxWithTimeout, s.infrastructures)

		s.in.Logger.SysLogger.Info(ctx, "Server exiting")
		exception <- isException
	}(ctx)

	select {
	case e := <-exception:
		if e {
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}

}

func (s *serverInterrupt) shutdown(ctx context.Context, services []*service) (isException bool) {
	errs := make(chan error, len(services))
	var wg sync.WaitGroup
	wg.Add(len(services))

	var do = func(ctx context.Context, name string, server IService) {
		defer wg.Done()
		if err := server.Shutdown(ctx); err != nil {
			s.in.Logger.SysLogger.Error(ctx, fmt.Errorf("shutdown %s error: %+v", name, err))
			errs <- err
		} else {
			s.in.Logger.SysLogger.Info(ctx, fmt.Sprintf("shutdown %s", name))
		}

	}

	for i := len(services); i > 0; i-- {
		serv := services[i-1]
		name := reflect.TypeOf(serv.IService).String()
		if serv.async {
			go do(ctx, name, serv)
		} else {
			do(ctx, name, serv)
		}

	}

	wg.Wait()

	close(errs)

	for err := range errs {
		if err != nil {
			isException = true
			s.in.Logger.SysLogger.Error(ctx, fmt.Errorf("shutdown error: %v", err))
		}
	}

	return isException
}
