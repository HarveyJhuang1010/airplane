package cron

import (
	"airplane/internal/components/apis"
	"airplane/internal/components/launcher"
	"airplane/internal/components/logger"
	"airplane/internal/config"
	"airplane/internal/infrastructure/database"
	"airplane/internal/infrastructure/rediscli"
	"airplane/internal/infrastructure/snowflake"
	"context"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"runtime/debug"
)

// Cmd portal server
var Cmd = &cobra.Command{
	Run:           runApplication,
	Use:           "cron",
	Short:         "Start cron server",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func runApplication(cmd *cobra.Command, args []string) {
	viper.Set("serverName", cmd.Use)

	binder := newBinder()
	if err := binder.Container.Invoke(run); err != nil {
		panic(err)
	}

	select {}
}

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
	CronService launcher.IService `name:"cron_srv_cron_service"`
	ApiService  *apis.Service     `name:"cron_srv_api_service"`
}

func run(in digIn) {
	ctx, cancel := context.WithCancel(context.Background())

	if bi, ok := debug.ReadBuildInfo(); ok {
		in.Logger.SysLogger.Info(ctx, "build info", zap.String("path", bi.Path))
		in.Logger.SysLogger.Info(ctx, "build info", zap.Any("setting", bi.Settings))
	}

	in.Server.Infrastructure(ctx, cancel,
		in.Redis,
		in.RDB,
	)

	// Background
	in.Server.Infrastructure(ctx, cancel,
		in.Snowflake, // 會用到 redis
	)

	in.Server.ListenAndServe(ctx, cancel,
		in.CronService,
		in.ApiService,
	)
}
