package apis

import (
	_ "airplane/docs"
	"airplane/internal/components/logger"
	"context"
	"errors"
	"fmt"
	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Apis struct {
	in dependence
}

func (r *Apis) New(ctx context.Context, opts ...serviceOption) *Service {
	service := &Service{
		in:      r.in,
		onStart: make([]func(), 0),
		Config: ServiceConfig{
			ListenAddress: "0.0.0.0",
			Port:          "9000",
			Trace:         false,
			AllowOrigins:  []string{"*"},
			AllowHeaders: []string{
				"Origin",
				"Authorization",
				"Access-Control-Allow-Origin",
				"Content-Type",
				"X-LOGIN-TOKEN"},
			ExposeHeaders: []string{
				"Content-Length",
				"Access-Control-Allow-Origin",
			},
			AllowMethods: []string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
				"PATCH",
				"OPTIONS",
			},
		},
		Logger: r.in.Logger.AppLogger,
	}

	for _, opt := range opts {
		opt(service)
	}
	service.engine = service.createService()

	return service
}

type Service struct {
	in dependence

	srv                 http.Server
	engine              *gin.Engine
	nativeRouterGroup   *gin.RouterGroup
	standardRouterGroup *gin.RouterGroup
	onStart             []func()

	Config ServiceConfig
	Logger logger.ILogger // 給業務邏輯用的 logger
}

func (s *Service) Run(ctx context.Context, stop context.CancelFunc) error {
	defer func() {
		if r := recover(); r != nil {
			s.in.Logger.SysLogger.Error(ctx, errors.New("panic"), zap.Any("panic", r))
		}
		stop()
	}()

	for _, handler := range s.onStart {
		handler()
	}

	s.srv = http.Server{
		ReadHeaderTimeout: time.Second * 10,
		Addr:              fmt.Sprintf("%s:%s", s.Config.ListenAddress, s.Config.Port),
		Handler:           s.engine,
	}

	s.in.Logger.SysLogger.Info(ctx, "Server is running", zap.String("address", s.srv.Addr))
	if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.in.Logger.SysLogger.Error(ctx, "Server Error", zap.Error(err))
	}

	return nil
}

func (s *Service) Shutdown(ctx context.Context) error {
	s.in.Logger.SysLogger.Info(nil, "start shutdown api server")
	if err := s.srv.Shutdown(ctx); err != nil {
		s.Logger.Error(ctx, "shutdown error", zap.Error(err))
	}
	s.in.Logger.SysLogger.Info(ctx, "end shutdown api server")

	return nil
}

func (s *Service) SetOnStarHandler(handler func()) {
	s.onStart = append(s.onStart, handler)
}

func (s *Service) GetEngine() *gin.Engine {
	return s.engine
}

func (s *Service) GetNativeRouterGroup() *gin.RouterGroup {
	return s.nativeRouterGroup
}

func (s *Service) GetStandardRouterGroup() *gin.RouterGroup {
	return s.standardRouterGroup
}

func (s *Service) createService() *gin.Engine {
	// Issue our HTTP Router
	e := gin.New()

	// for k8s health check
	e.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "health")
	})

	// RegisterService Handlers, Middleware, and API Modules
	s.registerMiddleware(e)
	s.nativeRouterGroup = s.registerNativeRoutes(e)
	s.standardRouterGroup = s.registerStandardRoutes(e)

	return e
}

func (s *Service) registerMiddleware(engine *gin.Engine) {
	engine.Use(
		gin.Recovery(),
	)

	engine.Use(ratelimit.RateLimiter(
		// storage
		ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
			Rate:  time.Second,
			Limit: 500,
		}),
		// options
		&ratelimit.Options{
			ErrorHandler: func(c *gin.Context, info ratelimit.Info) {
				c.String(http.StatusTooManyRequests, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
			},
			KeyFunc: func(c *gin.Context) string {
				key := c.GetHeader("Cf-Connecting-Ip")
				if len(key) > 0 {
					return key + c.Request.URL.Path
				}
				return c.ClientIP()
			},
		},
	))

	engine.Use(cors.New(cors.Config{
		AllowOrigins:  s.Config.AllowOrigins,
		AllowMethods:  s.Config.AllowMethods,
		AllowHeaders:  s.Config.AllowHeaders,
		ExposeHeaders: s.Config.ExposeHeaders,
	}))

	engine.Use(newRestAPIMiddleware(s.Logger, s.Config.Trace).Handler)

	engine.NoRoute(IgnoreStandardHandler, func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Not Found",
		})
	})
	engine.NoMethod(IgnoreStandardHandler, func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"code":    http.StatusMethodNotAllowed,
			"message": "Method Not Allowed",
		})
	})
}

func (s *Service) registerNativeRoutes(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/")
	group.Use(IgnoreStandardHandler)

	return group
}

func (s *Service) registerStandardRoutes(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/")

	return group
}
