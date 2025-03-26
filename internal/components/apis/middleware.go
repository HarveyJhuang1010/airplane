package apis

import (
	"airplane/internal/components/ctxs"
	"airplane/internal/components/errortool"
	"airplane/internal/components/logger"
	"airplane/internal/tools/timelogger"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

type IRestAPIMiddleware interface {
	Handler(ctx *gin.Context)
}

func newRestAPIMiddleware(log logger.ILogger, trace bool) IRestAPIMiddleware {

	return &restAPIMiddleware{
		trace:          trace,
		logger:         log,
		responseLogger: make(map[string]logger.ILogger),
	}
}

type restAPIMiddleware struct {
	trace  bool
	logger logger.ILogger

	responseLoggerMx sync.RWMutex
	responseLogger   map[string]logger.ILogger
}

func (mw *restAPIMiddleware) Handler(ctx *gin.Context) {
	if mw.trace {
		ctx.Set(timelogger.ContextKey, timelogger.NewTimeLogger())
	}

	ctxs.WithTraceID(ctx)
	traceID, _ := ctxs.Get[ctxs.TraceID](ctx)

	ctx.Next()

	if ignore, ok := ctxs.Get[ignoreStandardHandler](ctx); ok && ignore.Bool() {
		return
	}

	meta := &Meta{
		RequestID: traceID.String(),
	}

	if mw.trace {
		mw.decorateMeta(ctx, meta)
	}

	resp, err := ctx.MustGet(contextKey_FormatHandler).(IData).Format(ctx, meta)
	if err != nil {
		mw.logger.Error(ctx, err)
		ctx.Abort()
		return
	}

	mw.logResponse(ctx, meta, resp)

	for k, v := range resp.Headers {
		ctx.Header(k, v)
	}
	ctx.Data(resp.Status, resp.ContentType, resp.Data)
}

func (mw *restAPIMiddleware) decorateMeta(ctx *gin.Context, meta *Meta) {
	// Process time
	processTime, err := timelogger.GetTotalDuration(ctx)
	if err != nil {
		mw.logger.Error(ctx, err)
	}
	meta.ProcessTime = processTime.String()

	if processTime > 2*time.Second {
		timeLogs, err := timelogger.GetTimeLogs(ctx)
		if err != nil {
			mw.logger.Error(ctx, err)
		}
		meta.TimeLogs = timeLogs
	}
}

func (mw *restAPIMiddleware) getResponseLogger(name string) logger.ILogger {
	if name == "" {
		name = "defaultResponseLogger"
	}

	var getLoggerIfExists = func(name string) (logger.ILogger, bool) {
		mw.responseLoggerMx.RLock()
		defer mw.responseLoggerMx.RUnlock()

		if obj, ok := mw.responseLogger[name]; ok {
			return obj, ok
		}
		return nil, false
	}

	if obj, ok := getLoggerIfExists(name); ok {
		return obj
	}

	var createAndStoreLogger = func(name string) logger.ILogger {
		mw.responseLoggerMx.Lock()
		defer mw.responseLoggerMx.Unlock()

		if obj, ok := mw.responseLogger[name]; ok {
			return obj
		}

		obj := mw.logger.Named(name)
		mw.responseLogger[name] = obj
		return obj
	}

	return createAndStoreLogger(name)
}

func (mw *restAPIMiddleware) logResponse(ctx *gin.Context, meta *Meta, resp *Response) {
	if !mw.trace {
		return
	}

	var (
		respLogger = mw.getResponseLogger(resp.TraceNamed)
		msg        = fmt.Sprintf("response log [%s]", ctx.Request.URL)
	)

	switch respLogger.Level() {
	case "debug":
		respLogger.Info(ctx, msg,
			logger.Any("meta", meta),
			logger.Any("data", resp.Data))
	case "info":
		respLogger.Info(ctx, msg,
			logger.Any("meta", meta))
	default:

	}

	// error 都印
	// TODO: 之後做 handler 調整由外面決定內容，或什麼樣的 error 才印
	if resp.Error != nil {
		var errorStr string
		if e, ok := errortool.Parse(resp.Error); ok {
			errorStr = e.GetCode() + ": " + e.GetMessage()
		} else {
			errorStr = resp.Error.Error()
		}
		respLogger.Error(ctx, msg,
			logger.Any("meta", meta),
			logger.Any("data", resp.Data),
			logger.String("error", fmt.Sprintf("%+v", errorStr)))

		respLogger.Error(ctx, fmt.Sprintf("error:\n%v\n", resp.Error))

		return
	}

}
