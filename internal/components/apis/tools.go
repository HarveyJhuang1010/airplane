package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
)

// Convert 轉換不同 struct 的資料
func Convert[T any](ctx *gin.Context, from any) (to *T, err error) {
	defer func() {
		if e := recover(); e != nil {
			to = nil
			err = errors.Errorf("%s", e)
		}
	}()

	to = new(T)
	if err := copier.CopyWithOption(to, from, copier.Option{DeepCopy: true}); err != nil {
		return nil, err
	}

	return to, nil
}
