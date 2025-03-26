package portal

import (
	"airplane/internal/components/apis"
	"airplane/internal/components/launcher"
	"airplane/internal/components/logger"
	"airplane/internal/config"
	"airplane/internal/infrastructure/database"
	"airplane/internal/infrastructure/rediscli"
	"airplane/internal/infrastructure/snowflake"
	"context"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"runtime/debug"
)

type digIn struct {
	dig.In

	// tools
	Config *config.Config
	Logger *logger.Loggers
	Server *launcher.Launcher
	// infrastructure
	Redis     *rediscli.Redis
	RDB       *database.DB `name:"dbMaster"`
	Snowflake *snowflake.Snowflake
	// server
	ApiService *apis.Service `name:"portal_api_service"`
}

func run(in digIn) {
	ctx, cancel := context.WithCancel(context.Background())

	if bi, ok := debug.ReadBuildInfo(); ok {
		in.Logger.SysLogger.Info(ctx, "build info", zap.String("path", bi.Path))
		in.Logger.SysLogger.Info(ctx, "build info", zap.Any("setting", bi.Settings))
	}

	// Infrastructure
	in.Server.Infrastructure(ctx, cancel,
		in.Redis,
		in.RDB,
	)

	// Background
	in.Server.Infrastructure(ctx, cancel,
		in.Snowflake,
	)

	in.Server.ListenAndServe(ctx, cancel, in.ApiService)
}
