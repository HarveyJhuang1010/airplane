package logger

import (
	"context"
	"go.uber.org/zap"
)

type mockLogger struct {
}

func (m *mockLogger) Level() string {
	return ""
}

func (m *mockLogger) Named(s string) ILogger {
	return &mockLogger{}
}

func (m *mockLogger) WithOptions(opts ...zap.Option) ILogger {
	return &mockLogger{}
}

func (m *mockLogger) Sugar() *SugaredLogger {
	return &SugaredLogger{}
}

func (m *mockLogger) Debug(ctx context.Context, message string, fields ...Field) {}

func (m *mockLogger) Info(ctx context.Context, message string, fields ...Field) {}

func (m *mockLogger) Warn(ctx context.Context, message string, fields ...Field) {}

func (m *mockLogger) Error(ctx context.Context, err any, fields ...Field) {}

func (m *mockLogger) Panic(ctx context.Context, err any, fields ...Field) {}
