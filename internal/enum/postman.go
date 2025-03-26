package enum

//go:generate enumer -type=PostmanContactType -trimprefix=PostmanContactType -yaml -json -text -transform=snake --output=zzz_enumer_postmanContactType.go
type PostmanContactType int32

const (
	PostmanContactTypeUnknown PostmanContactType = iota
	PostmanContactTypeEmail
	PostmanContactTypePhone
)
