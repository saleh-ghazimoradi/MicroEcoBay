package cmd

import (
	"fmt"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/infra/db"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/slg"

	"github.com/spf13/cobra"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("down called")
		d, err := db.PostDBConnection(db.PostDBMigrateDrop)
		if err != nil {
			slg.Logger.Error("failed to connect database", "error", err)
			return
		}

		if err := db.PostDBMigrateDrop(d); err != nil {
			slg.Logger.Error("failed to migrate drop", "error", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
}
