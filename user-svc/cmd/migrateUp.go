package cmd

import (
	"fmt"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/config"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/infra/postgresql"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/domain"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/logger"
	"log"

	"github.com/spf13/cobra"
)

// migrateUpCmd represents the migrateUp command
var migrateUpCmd = &cobra.Command{
	Use:   "migrateUp",
	Short: "It runs the migrate up command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrateUp called")

		cfg, err := config.GetConfig()
		if err != nil {
			log.Fatalf("failed to get config: %v", err)
		}

		log := logger.NewLogger()

		postDB := postgresql.NewPostgresql(
			postgresql.WithHost(cfg.Postgresql.Host),
			postgresql.WithPort(cfg.Postgresql.Port),
			postgresql.WithUser(cfg.Postgresql.User),
			postgresql.WithPassword(cfg.Postgresql.Password),
			postgresql.WithName(cfg.Postgresql.Name),
			postgresql.WithMaxOpenConn(cfg.Postgresql.MaxOpenConn),
			postgresql.WithMaxIdleConn(cfg.Postgresql.MaxIdleConn),
			postgresql.WithMaxIdleTime(cfg.Postgresql.MaxIdleTime),
			postgresql.WithSSLMode(cfg.Postgresql.SSLMode),
			postgresql.WithTimeout(cfg.Postgresql.Timeout),
			postgresql.WithLogger(&log),
		)

		gormDB, err := postDB.Connect()
		if err != nil {
			log.Fatal().Err(err).Msg("error connecting to database")
		}

		if err := gormDB.Migrator().AutoMigrate(&domain.User{}, &domain.Address{}); err != nil {
			log.Fatal().Err(err).Msg("error migrating database")
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateUpCmd)
}
