package logger

import (
	config "airplane/internal/config"
	"context"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type SugaredLogger = zap.SugaredLogger
type Option = zap.Option

type ILogger interface {
	Level() string
	Named(s string) ILogger
	WithOptions(opts ...zap.Option) ILogger
	Sugar() *SugaredLogger
	Debug(ctx context.Context, message string, fields ...Field)
	Info(ctx context.Context, message string, fields ...Field)
	Warn(ctx context.Context, message string, fields ...Field)
	Error(ctx context.Context, err any, fields ...Field)
	Panic(ctx context.Context, err any, fields ...Field)
}

func Mock() *Loggers {
	return &Loggers{
		SysLogger: &mockLogger{},
		AppLogger: &mockLogger{},
	}
}

func New(in digIn) *Loggers {
	loggers := &Loggers{
		in: in,
	}
	return loggers.initialize()
}

type digIn struct {
	dig.In

	AppConf *config.Config
}

type Loggers struct {
	in digIn

	SysLogger ILogger
	AppLogger ILogger
}

func (l *Loggers) New(config Config) ILogger {
	return newLogger(l.in, config)
}

func (l *Loggers) initialize() *Loggers {
	l.SysLogger = l.initializeSysLogger()
	l.AppLogger = l.initializeAppLogger()
	return l
}

func (l *Loggers) initializeSysLogger() ILogger {
	return newLogger(l.in, Config{
		Level:    l.in.AppConf.Logger.SysLogger,
		Category: "SysLogger",
	})
}

func (l *Loggers) initializeAppLogger() ILogger {
	return newLogger(l.in, Config{
		Level:    l.in.AppConf.Logger.AppLogger,
		Category: "AppLogger",
	})
}

func getZapLevel(l string) zapcore.Level {
	switch l {
	case zapcore.DebugLevel.String(): // "debug"
		return zapcore.DebugLevel
	case zapcore.InfoLevel.String(): // "info"
		return zapcore.InfoLevel
	case zapcore.WarnLevel.String(): // "warn"
		return zapcore.WarnLevel
	case zapcore.ErrorLevel.String(): // "error"
		return zapcore.ErrorLevel
	case zapcore.DPanicLevel.String(): // "dpanic"
		return zapcore.DPanicLevel
	case zapcore.PanicLevel.String(): // "panic"
		return zapcore.PanicLevel
	case zapcore.FatalLevel.String(): // "fatal"
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
