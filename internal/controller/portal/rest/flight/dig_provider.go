package flight

import (
	"airplane/internal/components/apis"
	"airplane/internal/components/logger"
	"airplane/internal/controller/portal/rest/common"
	"airplane/internal/core/usecase/flight"
	"go.uber.org/dig"
)

func New(in digIn) digOut {
	dep := dependence{
		digIn: in,
	}
	dep.ApiV1Data = func() *common.ApiV1Data {
		return common.V1Data.New(common.V1DataOpt.WithTraceNamed("userApiV1"))
	}
	dep.StandardData = func() *apis.StandardData {
		return common.StandardData.New(common.StandardDataOtp.WithTraceNamed("userApiV1"))
	}
	dep.RawData = func() *apis.RawData {
		return common.RawData.New(common.RawDataOtp.WithTraceNamed("userApiV1"))
	}
	dep.ErrorMapping = common.NewErrorRepository()

	return digOut{
		Controller: newController(dep),
		Middleware: newMiddleware(dep),
	}
}

type dependence struct {
	digIn

	ApiV1Data    func() *common.ApiV1Data
	StandardData func() *apis.StandardData
	RawData      func() *apis.RawData
	ErrorMapping *common.ErrorMappingRepository
}

type digIn struct {
	dig.In

	Logger   *logger.Loggers
	Response apis.IResponse

	// core
	Flight *flight.Usecase
}

type digOut struct {
	dig.Out

	Controller *Controller
	Middleware *Middleware
}

func newController(in dependence) *Controller {
	return &Controller{
		Flight: newFlight(in),
	}
}

type Controller struct {
	Flight *Flight
}

func newMiddleware(dependence dependence) *Middleware {
	return &Middleware{}
}

type Middleware struct {
}
