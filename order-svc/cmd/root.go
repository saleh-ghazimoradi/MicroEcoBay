package cmd

import (
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/config"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/slg"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "order_service",
	Short: "order_service is a part of the MicroEcoBay microservices ecosystem",
}

func Execute() {
	err := os.Setenv("TZ", time.UTC.String())
	if err != nil {
		panic(err)
	}

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	err := config.LoadConfig()
	if err != nil {
		slg.Logger.Error("there went something wrong while loading config file", "error", err)
	}
}
