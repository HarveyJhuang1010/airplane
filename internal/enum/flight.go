package enum

//go:generate enumer -type=FlightStatus -trimprefix=FlightStatus -yaml -json -text -transform=snake --output=zzz_enumer_FlightStatus.go
type FlightStatus int32

const (
	FlightStatusScheduled FlightStatus = iota
	FlightStatusBoarding
	FlightStatusDeparted
	FlightStatusArrived
	FlightStatusCancelled
)
