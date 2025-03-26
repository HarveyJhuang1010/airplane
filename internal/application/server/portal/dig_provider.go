package portal

import (
	"airplane/internal/components/apis"
	"airplane/internal/components/logger"
	"airplane/internal/config"
	"go.uber.org/dig"
)

func NewServer(in digIn) digOut {
	return digOut{
		ApiService: newService(in, "portal_api_service"),
	}
}

type digIn struct {
	dig.In

	Logger *logger.Loggers
	Config *config.Config
	Apis   *apis.Apis
}

type digOut struct {
	dig.Out

	ApiService *apis.Service `name:"portal_api_service"`
}
