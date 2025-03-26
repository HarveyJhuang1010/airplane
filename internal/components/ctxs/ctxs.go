package ctxs

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"reflect"
)

// 亂數避免外部 key 衝突
const key = "_FWs8I"

// Get value from context
func Get[T any](ctx context.Context) (T, bool) {
	var zero T
	typeName := reflect.TypeOf(zero).Name() + key

	v := ctx.Value(typeName)
	if v == nil {
		return zero, false
	}
	return v.(T), true
}

// Set value to context
func Set[T any](ctx context.Context, val T) context.Context {
	typeName := reflect.TypeOf(val).Name() + key

	switch c := ctx.(type) {
	case *gin.Context:
		c.Set(typeName, val)
	default:
		ctx = context.WithValue(ctx, typeName, val)
	}

	return ctx
}

type TraceID struct {
	string
}

func (t TraceID) String() string {
	return t.string
}

func WithTraceID(ctx context.Context) context.Context {
	var uuid = TraceID{uuid.New().String()}
	return Set(ctx, uuid)
}
