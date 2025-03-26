package enum

//go:generate enumer -type=TrailOperatorType -trimprefix=TrailOperatorType -yaml -json -text -transform=snake --output=zzz_enumer_trailOperatorType.go
type TrailOperatorType int32

const (
	TrailOperatorTypeUnknown TrailOperatorType = iota
	TrailOperatorTypeSystem
)
