package portal

import (
	"airplane/internal/controller/portal/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Cmd portal all server
var Cmd = &cobra.Command{
	Run:           runAllApplication,
	Use:           "portal",
	Short:         "Start portal server",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func runAllApplication(cmd *cobra.Command, args []string) {
	viper.Set("serverName", cmd.Use)

	var (
		routers = []interface{}{
			rest.NewBookingRouter,
			rest.NewFlightRouter,
			rest.NewPaymentRouter,
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
