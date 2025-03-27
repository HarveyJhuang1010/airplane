package enum

//go:generate enumer -type=BookingStatus -trimprefix=BookingStatus -yaml -json -text -sql -transform=snake --output=zzz_enumer_BookingStatus.go
type BookingStatus int32

const (
	BookingStatusPending BookingStatus = iota
	BookingStatusConfirming
	BookingStatusConfirmed
	BookingStatusCanceling
	BookingStatusCancelled
	BookingStatusExpired
	BookingStatusOverbooked
)
