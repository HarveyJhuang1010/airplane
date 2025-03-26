package logger

import (
	"airplane/internal/components/ctxs"
	"airplane/internal/config"
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

// func newLogger(level zapcore.Level, serviceName string, category string, env string) *logger {
func newLogger(in digIn, cfg Config) *logger {
	level := getZapLevel(cfg.Level)
	configEncoding := "json"
	encoderConfigLevelEncoder := zapcore.CapitalLevelEncoder
	if in.AppConf.Env == config.EnvLocalTag {
		configEncoding = "console"
		encoderConfigLevelEncoder = zapcore.CapitalColorLevelEncoder
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encoderConfigLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(level), // 日志级别
		DisableStacktrace: true,
		Development:       false,          // 开发模式，堆栈跟踪
		Encoding:          configEncoding, // 输出格式 console 或 json
		EncoderConfig:     encoderConfig,  // 编码器配置
		InitialFields: map[string]interface{}{
			"version":  "1",
			"category": cfg.Category,
			"service":  in.AppConf.GetServerName(),
		}, // 初始化字段，如：添加一个服务器名称
		OutputPaths:      []string{"stdout"}, // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"},
	}

	zapLogger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	if cfg.Named != "" {
		zapLogger = zapLogger.Named(cfg.Named)
	}

	return &logger{
		in:     in,
		cfg:    cfg,
		logger: zapLogger,
	}
}

type Error interface {
	string | any
}

type logger struct {
	in     digIn
	cfg    Config
	logger *zap.Logger
}

func (lg *logger) Level() string {
	return lg.logger.Level().String()
}

func (lg *logger) Named(s string) ILogger {

	if val, ok := lg.in.AppConf.Logger.Named[strings.ToLower(s)]; ok {
		named := ""
		if lg.cfg.Named == "" {
			named = s
		} else {
			named = strings.Join([]string{lg.cfg.Named, s}, ".")
		}
		return newLogger(lg.in, Config{
			Level:    val,
			Category: lg.cfg.Category,
			Named:    named,
		})
	}

	return &logger{
		in:     lg.in,
		logger: lg.logger.Named(s),
	}
}

func (lg *logger) WithOptions(opts ...zap.Option) ILogger {
	return &logger{
		in:     lg.in,
		logger: lg.logger.WithOptions(opts...),
	}
}

func (lg *logger) Sugar() *SugaredLogger {
	return lg.logger.Sugar()
}

func (lg *logger) Debug(ctx context.Context, message string, fields ...Field) {
	lg.logger.Debug(message, enrichFields(ctx, fields...)...)
	lg.logger.Sugar()
}

func (lg *logger) Debugf(ctx context.Context, message string, fields ...Field) {
	lg.logger.Info(message, enrichFields(ctx, fields...)...)
}

func (lg *logger) Info(ctx context.Context, message string, fields ...Field) {
	lg.logger.Info(message, enrichFields(ctx, fields...)...)
}

func (lg *logger) Warn(ctx context.Context, message string, fields ...Field) {
	lg.logger.Warn(message, enrichFields(ctx, fields...)...)
}

func (lg *logger) Error(ctx context.Context, err any, fields ...Field) {
	lg.logger.Error(fmt.Sprintf("%+v", err), enrichFields(ctx, fields...)...)
}

func (lg *logger) Panic(ctx context.Context, err any, fields ...Field) {
	lg.logger.Panic(fmt.Sprintf("%+v", err), enrichFields(ctx, fields...)...)
}

func enrichFields(ctx context.Context, fields ...Field) []Field {
	return append(fields, traceID(ctx))
}

func traceID(ctx context.Context) Field {
	if ctx == nil {
		return zap.String("traceID", "")
	}

	var chainID string

	if val, ok := ctxs.Get[ctxs.TraceID](ctx); ok {
		chainID = val.String()
	}

	return zap.String("traceID", chainID)
}
