package enum

//go:generate enumer -type=BookingStatus -trimprefix=BookingStatus -yaml -json -text -transform=snake --output=zzz_enumer_BookingStatus.go
type BookingStatus int32

const (
	BookingStatusPending BookingStatus = iota
	BookingStatusConfirmed
	BookingStatusCancelled
	BookingStatusExpired
	BookingStatusOverbooked
	BookingStatusConfirming
)
