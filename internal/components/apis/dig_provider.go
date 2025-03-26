package apis

import (
	"airplane/internal/components/logger"
	"go.uber.org/dig"
)

func New(in digIn) digOut {
	dep := dependence{
		digIn: in,
	}

	self := &digProvider{
		in: dep,
		out: digOut{
			Response: responseHandler,
			Apis:     &Apis{in: dep},
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
}

type digIn struct {
	dig.In

	Logger *logger.Loggers
}

type digOut struct {
	dig.Out

	Response IResponse
	Apis     *Apis
}
