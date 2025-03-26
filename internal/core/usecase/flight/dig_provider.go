package flight

import (
	"airplane/internal/core/repositories/rdb"
	"go.uber.org/dig"
)

func New(in digIn) digOut {
	dep := dependence{
		digIn: in,
	}

	return digOut{
		Usecase: newFlightUsecase(dep),
	}
}

type dependence struct {
	digIn
}

type digIn struct {
	dig.In

	DBRepository *rdb.Repository
}

type digOut struct {
	dig.Out

	Usecase *Usecase
}

func newFlightUsecase(in dependence) *Usecase {
	return &Usecase{
		ListFlight: newListFlight(in),
	}
}

type Usecase struct {
	ListFlight *ListFlight
}
