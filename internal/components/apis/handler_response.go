package apis

import (
	"github.com/gin-gonic/gin"
)

type IResponse interface {
	Data(ctx *gin.Context, f IData)
}

var responseHandler = &responseHandlerImpl{}

type responseHandlerImpl struct {
}

func (r *responseHandlerImpl) Data(ctx *gin.Context, f IData) {
	ctx.Set(contextKey_FormatHandler, f)
}
