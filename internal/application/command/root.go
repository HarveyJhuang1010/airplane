package cmd

import (
	"fmt"
	"os"

	"airplane/internal/application/command/cron"
	"airplane/internal/application/command/portal"
	"airplane/internal/application/command/qworker"
	"github.com/spf13/cobra"
)

// Root command
var (
	rootCmd = &cobra.Command{
		SilenceUsage: true,
	}
)

func Execute() {
	rootCmd.AddCommand(
		cron.Cmd,
		portal.Cmd,
		portal.BookingCmd,
		portal.PaymentCmd,
		portal.FlightCmd,
		qworker.Cmd,
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		os.Exit(1)
	}
}
