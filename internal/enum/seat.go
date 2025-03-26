package enum

//go:generate enumer -type=SeatStatus -trimprefix=SeatStatus -yaml -json -text -sql -transform=snake --output=zzz_enumer_SeatStatus.go
type SeatStatus int32

const (
	SeatStatusAvailable SeatStatus = iota
	SeatStatusHeld
	SeatStatusBooked
)
