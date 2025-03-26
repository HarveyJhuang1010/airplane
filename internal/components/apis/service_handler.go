package apis

import (
	_ "airplane/docs"
	"airplane/internal/components/ctxs"
	"context"
	"github.com/gin-gonic/gin"
)

type ignoreStandardHandler struct {
	b bool
}

func (s ignoreStandardHandler) Bool() bool {
	return s.b
}

func withIgnoreStandardHandler(ctx context.Context) context.Context {
	var val = ignoreStandardHandler{true}
	return ctxs.Set(ctx, val)
}

func IgnoreStandardHandler(ctx *gin.Context) {
	withIgnoreStandardHandler(ctx)
	ctx.Next()
}
