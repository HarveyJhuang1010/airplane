package enum

//go:generate enumer -type=CabinClassCode -trimprefix=CabinClassCode -yaml -json -text -sql -transform=snake --output=zzz_enumer_CabinClassCode.go
type CabinClassCode int32

const (
	CabinClassCodeEconomyStandard CabinClassCode = iota
	CabinClassCodeEconomyFlex
	CabinClassCodeBusinessBasic
	CabinClassCodeBusinessStandard
)
