package bo

import (
	"time"

	"airplane/internal/enum"
)

type Postman2faUserSendVerifyCodeCond struct {
	Type    enum.PostmanContactType
	Contact string // 聯絡方式

	Code         string // 驗證碼
	CodeLifeTime string // 驗證碼有效時間
}

type PostmanNotifyUserLoginSendCond struct {
	Type    enum.PostmanContactType
	Contact string // 聯絡方式

	Account  string
	Device   string
	IP       string
	Browser  string
	Location string
	Time     time.Time
}

type PostmanAdminCreateCond struct {
	Type    enum.PostmanContactType
	Contact string // 聯絡方式

	Password string // 驗證碼
}

type PostmanAdminResetPasswordCond struct {
	Type    enum.PostmanContactType
	Contact string // 聯絡方式

	Password string // 驗證碼
}

type PostmanNotifyUpdatePasswordSendCond struct {
	Contact string // 聯絡方式
	Type    enum.PostmanContactType
	Time    time.Time
}

type PostmanNotifySignApplySendCond struct {
	Contact string // 聯絡方式
	Type    enum.PostmanContactType

	ApplyAdminEmail string
	ApplyType       string
	ApplyTime       time.Time
}

type PostmanNotifySignResultSendCond struct {
	Contact string // 聯絡方式
	Type    enum.PostmanContactType

	SignAdminEmail string
	ApplyType      string
	SignResult     string
	SignedTime     time.Time
}
