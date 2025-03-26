package enum

//go:generate enumer -type=DeviceType -trimprefix=DeviceType -yaml -json -text -transform=snake --output=zzz_enumer_deviceType.go
type DeviceType int32

const (
	DeviceTypeUnknown DeviceType = iota
	DeviceTypeMobile
	DeviceTypeTablet
	DeviceTypeDesktop
	DeviceTypeBot
)
