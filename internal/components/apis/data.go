package apis

import (
	"github.com/gin-gonic/gin"
)

type IData interface {
	Format(ctx *gin.Context, meta *Meta) (*Response, error)
}

// 檢查是否符合 IData 介面
var _ IData = (*StandardData)(nil)
var _ IData = (*RawData)(nil)
