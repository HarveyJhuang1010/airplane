package qworker

import (
	"airplane/internal/components/apis"
	"context"
	"strconv"
)

func newApiService(in digIn) *apis.Service {

	svc := in.Apis.New(
		context.Background(),
		apis.ServiceOption().WithLogger(in.Logger.SysLogger.Named("apiService")),
		apis.ServiceOption().WithConfig(apis.ServiceConfig{
			ListenAddress: in.Config.Rest.ListenAddress,
			Port:          strconv.Itoa(in.Config.Rest.Port),
			Trace:         in.Config.Rest.Trace,
			AllowOrigins:  in.Config.Rest.AllowOrigins,
			AllowHeaders:  in.Config.Rest.AllowHeaders,
			ExposeHeaders: in.Config.Rest.ExposeHeaders,
			AllowMethods:  in.Config.Rest.AllowMethods,
		}),
	)

	return svc
}
