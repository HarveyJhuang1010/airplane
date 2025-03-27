package enum

//go:generate enumer -type=UserStatus -trimprefix=UserStatus -yaml -json -text -sql -transform=snake --output=zzz_enumer_UserStatus.go
type UserStatus int32

const (
	UserStatusUnverified UserStatus = iota // 未啟用
	UserStatusEnable                       // 已啟用
	UserStatusDisable                      // 已停用
	UserStatusFrozen                       // 凍結中
)
