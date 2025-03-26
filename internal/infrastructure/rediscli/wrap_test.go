package rediscli

import (
	"airplane/internal/config"
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"airplane/internal/appcontext"
	"airplane/internal/errs"
	"airplane/internal/logging"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

type redisCliTestApp struct {
	dig.In

	RedisCli *Redis
}

type redisCliTestSuite struct {
	suite.Suite
	app    *redisCliTestApp
	logger *zap.Logger
	cfg    *config.Config
	ctx    *appcontext.AppContext
}

func (s *redisCliTestSuite) SetupSuite() {
	// Init Connection & Config
	w := httptest.NewRecorder()
	gctx, _ := gin.CreateTestContext(w)
	appctx := appcontext.New()
	appctx.SetApiContext(gctx)
	cfg := config.New()
	logger := logging.NewZapLogger("bgx-test")
	appcontext.SetDefaultLogger(logger)
	appctx.SetApiLogger(logger)
	Initialize(cfg.Redis)

	s.logger = logger
	s.cfg = cfg
	s.ctx = appctx

	// Init Your Provider
	binder := dig.New()
	s.Require().Nil(binder.Provide(NewCacheClient))
	s.Require().Nil(binder.Invoke(func(app redisCliTestApp) {
		s.app = &app
	}))
}

func (s *redisCliTestSuite) SetupTest() {
}

func (s *redisCliTestSuite) TearDownTest() {

}

func (s *redisCliTestSuite) TearDownSuite() {
	Finalize()
}

func (s *redisCliTestSuite) TestKeyValue() {

	s.Run("GetKey", func() {
		_, err := s.app.RedisCli.Get(s.ctx, "key_not_exist").Result()
		s.Require().False(errors.Is(err, errs.ErrRedisNotFound))
	})
	s.Run("SetKey", func() {
		_, err := s.app.RedisCli.Set(s.ctx, "key", "value", time.Duration(0)).Result()
		s.Require().Nil(err)
	})
	s.Run("GetKey", func() {
		val, err := s.app.RedisCli.Get(s.ctx, "key").Result()
		s.Require().Nil(err)
		s.Require().Equal("value", val, "value not match")
	})
	s.Run("Check Wrap", func() {
		_, err := Wrap(s.app.RedisCli.Get(s.ctx, "key_not_exist")).Result()
		s.Require().True(errors.Is(err, errs.ErrRedisNotFound))
	})

	s.Run("zscore", func() {
		_, err := Wrap(s.app.RedisCli.ZScore(s.ctx, "key_not_exist", "test")).Result()
		s.Require().True(errors.Is(err, errs.ErrRedisNotFound))
	})
}

func (s *redisCliTestSuite) TestWrapError() {

	s.Run("operation failed error with wrap", func() {
		_, err := Wrap(s.app.RedisCli.LIndex(s.ctx, "key", 0)).Result()
		s.Require().True(errors.Is(err, errs.ErrRedisOperationFailed))
	})
	s.Run("operation failed error without wrap", func() {
		_, err := s.app.RedisCli.LIndex(s.ctx, "key", 0).Result()
		s.Require().False(errors.Is(err, errs.ErrRedisOperationFailed))
	})
}

func TestRedisCli(t *testing.T) {
	suite.Run(t, &redisCliTestSuite{})
}
