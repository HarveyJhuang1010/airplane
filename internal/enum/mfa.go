package enum

//go:generate enumer -type=MfaAction -trimprefix=MfaAction -yaml -json -text -transform=snake --output=zzz_enumer_mfaAction.go
type MfaAction uint32

const (
	MfaActionUnknown MfaAction = iota + 1
)

//go:generate enumer -type=MfaType -trimprefix=MfaType -yaml -json -text -transform=snake --output=zzz_enumer_mfaType.go
type MfaType uint32

const (
	MfaTypeEmail MfaType = 1 << iota
	MfaTypeSms
	MfaTypeGoogleAuthenticator
)
