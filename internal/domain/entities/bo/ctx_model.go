package bo

import (
	"time"
)

type CtxRequestInfo struct {
	RequestTime   time.Time // 請求當下時間
	Host          string
	ClientIP      string
	UserAgent     string
	XForwardedFor string
}

type CtxUserSession struct {
	UserID int64
}
