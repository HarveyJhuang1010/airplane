package timelogger

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const ContextKey = "timeLogger"

var (
	err_TimeLoggerTypeError = errors.New("必須包含 *timeLogger 進 context 才能正常使用 TimeLog")
)

func NewTimeLogger() *timeLogger {
	return &timeLogger{
		startTime: time.Now(),
	}
}

func LogTime(ctx context.Context) func() {
	timeLogger, existed := ctx.Value(ContextKey).(*timeLogger)
	if !existed {
		return func() {}
	}

	return timeLogger.logFuncTime(getFuncName())
}

func GetTotalDuration(ctx context.Context) (time.Duration, error) {
	timeLogger, existed := ctx.Value(ContextKey).(*timeLogger)
	if !existed {
		return time.Duration(0), err_TimeLoggerTypeError
	}
	return timeLogger.getTotalDuration(), nil
}

func GetTimeLogs(ctx context.Context) ([]string, error) {
	timeLogger, existed := ctx.Value(ContextKey).(*timeLogger)
	if !existed {
		return nil, err_TimeLoggerTypeError
	}

	logs := make([]string, 0, len(timeLogger.process))
	for idx, _ := range timeLogger.process {
		execTime := timeLogger.getDuration(idx)
		logs = append(logs, execTime.String())
	}

	return logs, nil
}

type timeLogger struct {
	mux sync.RWMutex

	startTime time.Time
	process   []execTime
}

func (logger *timeLogger) getTotalDuration() time.Duration {
	return time.Since(logger.startTime)
}

func (logger *timeLogger) logFuncTime(funcName string) func() {
	now := time.Now()

	logger.mux.Lock()
	logger.process = append(logger.process, execTime{FuncName: funcName})
	idx := len(logger.process) - 1
	logger.mux.Unlock()

	return func() {
		logger.mux.Lock()
		logger.process[idx].ExecutionTime = time.Since(now)
		logger.mux.Unlock()
	}
}

func (logger *timeLogger) getDuration(idx int) execTime {
	logger.mux.RLock()
	defer logger.mux.RUnlock()

	return logger.process[idx]
}

func getFuncName() string {
	pc, _, line, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name() + ":" + fmt.Sprint(line)
	return funcName[strings.LastIndex(funcName, "/")+1:]
}

type execTime struct {
	FuncName      string
	ExecutionTime time.Duration // 這個函數執行的時間
}

func (e execTime) String() string {
	return fmt.Sprintf("+%-12s %s", e.ExecutionTime, e.FuncName)
}
