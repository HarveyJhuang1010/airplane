package user

import (
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/infrastructure/snowflake"
	"go.uber.org/dig"
)

func New(in digIn) digOut {
	dep := dependence{
		digIn: in,
	}

	return digOut{
		Usecase: newUserUsecase(dep),
	}
}

type dependence struct {
	digIn
}

type digIn struct {
	dig.In

	Snowflake    *snowflake.Snowflake
	DBRepository *rdb.Repository
}

type digOut struct {
	dig.Out

	Usecase *Usecase
}

func newUserUsecase(in dependence) *Usecase {
	return &Usecase{
		GetUser: newGetUser(in),
	}
}

type Usecase struct {
	GetUser *GetUser
}
