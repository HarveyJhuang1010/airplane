package enum

//go:generate enumer -type=UserLoginType -trimprefix=UserLoginType -yaml -json -text -transform=snake --output=zzz_enumer_userLoginType.go
type UserLoginType int32

const (
	UserLoginTypeUnknown UserLoginType = iota
	UserLoginTypeEmail
	UserLoginTypePhone
	UserLoginTypeGoogle
	UserLoginTypeFacebook
	UserLoginTypeApple
	UserLoginTypeLine
)

const (
	UserLoginTypeDealer UserLoginType = 101
)

//go:generate enumer -type=UserLoginStatus -trimprefix=UserLoginStatus -yaml -json -text -transform=snake --output=zzz_enumer_userLoginStatus.go
type UserLoginStatus int32

const (
	UserLoginStatusDisable UserLoginStatus = iota
	UserLoginStatusEnable
)

//go:generate enumer -type=UserLoginSessionStep -trimprefix=UserLoginSessionStep -yaml -json -text -transform=snake --output=zzz_enumer_userLoginSessionStep.go
type UserLoginSessionStep int32

const (
	UserLoginSessionStepIdentify    UserLoginSessionStep = iota + 1 // 識別身份
	UserLoginSessionStepOtpResend                                   // 重發驗證碼
	UserLoginSessionStepOtpVerified                                 // 驗證成功
	UserLoginSessionStepFinish                                      // 登入成功
)
