package portal

import (
	"airplane/internal/controller/portal/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Cmd portal all server
var FlightCmd = &cobra.Command{
	Run:           runFlightApplication,
	Use:           "flight",
	Short:         "Start flight portal server",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func runFlightApplication(cmd *cobra.Command, args []string) {
	viper.Set("serverName", cmd.Use)

	var (
		routers = []interface{}{
			rest.NewFlightRouter,
		}
	)

	binder := newBinder()

	for _, r := range routers {
		if err := binder.Container.Invoke(r); err != nil {
			panic(err)
		}
	}
	if err := binder.Container.Invoke(run); err != nil {
		panic(err)
	}

	select {}
}
