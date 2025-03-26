package enum

//go:generate enumer -type=IdentityType -trimprefix=IdentityType -yaml -json -text -transform=snake --output=zzz_enumer_identityType.go
type IdentityType int32

const (
	IdentityTypeUnknown IdentityType = iota + 1 // 管理員 for 後台
)
