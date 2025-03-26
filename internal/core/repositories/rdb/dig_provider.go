package rdb

import (
	"airplane/internal/infrastructure/database"
	"go.uber.org/dig"
)

func New(in digIn) digOut {
	dep := dependence{
		digIn: in,
	}

	self := &digProvider{
		in: dep,
		out: digOut{
			Repository: newRepository(dep),
		},
	}

	return self.out
}

type digProvider struct {
	in dependence

	out digOut
}

type dependence struct {
	digIn

	common common
}

type digIn struct {
	dig.In

	DBMaster *database.DB `name:"dbMaster"`
}

type digOut struct {
	dig.Out

	Repository *Repository
}
