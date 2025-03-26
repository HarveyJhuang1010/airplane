package portal

import (
	"airplane/internal/controller/portal/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Cmd portal all server
var BookingCmd = &cobra.Command{
	Run:           runBookingApplication,
	Use:           "booking",
	Short:         "Start booking portal server",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func runBookingApplication(cmd *cobra.Command, args []string) {
	viper.Set("serverName", cmd.Use)

	var (
		routers = []interface{}{
			rest.NewBookingRouter,
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
