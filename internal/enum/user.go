package enum

//go:generate enumer -type=UserStatus -trimprefix=UserStatus -yaml -json -text -transform=snake --output=zzz_enumer_userStatus.go
type UserStatus int32

const (
	UserStatusUnverified UserStatus = iota // 未啟用
	UserStatusEnable                       // 已啟用
	UserStatusDisable                      // 已停用
	UserStatusFrozen                       // 凍結中
)

//go:generate enumer -type=UserGender -trimprefix=UserGender -yaml -json -text -transform=snake --output=zzz_enumer_userGender.go
type UserGender int32

const (
	UserGenderUnknown UserGender = iota
	UserGenderMale
	UserGenderFemale
)
