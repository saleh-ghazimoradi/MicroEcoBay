package cmd

import (
	"fmt"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/gateway"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/slg"
	"os"

	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "to run user service http connection",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("http called")
		if err := gateway.Server(); err != nil {
			slg.Logger.Error("the server encountered an error: " + err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
