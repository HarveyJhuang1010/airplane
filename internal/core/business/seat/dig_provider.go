package seat

import (
	"airplane/internal/components/logger"
	"airplane/internal/infrastructure/mqcli"
	"go.uber.org/dig"
)

func New(in digIn) digOut {
	dep := dependence{
		digIn: in,
	}

	return digOut{
		Usecase: newSeatUsecase(dep),
	}
}

type dependence struct {
	digIn
}

type digIn struct {
	dig.In

	Logger *logger.Loggers
	Kafka  *mqcli.Kafka
}

type digOut struct {
	dig.Out

	Usecase *Usecase
}

func newSeatUsecase(in dependence) *Usecase {
	return &Usecase{
		BookSeat:    newBookSeat(in),
		ReleaseSeat: newReleaseSeat(in),
	}
}

type Usecase struct {
	BookSeat    *BookSeat
	ReleaseSeat *ReleaseSeat
}
