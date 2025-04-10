package cmd

import (
	"github.com/saleh-ghazimoradi/MicroEcoBay/payment_service/config"
	"github.com/saleh-ghazimoradi/MicroEcoBay/payment_service/slg"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "payment_service",
	Short: "A brief description of your application",
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
