package cmd

import (
	"fmt"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/config"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/infra/postgresql"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/gateway/rest/handlers"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/gateway/rest/middlewares"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/gateway/rest/routes"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/logger"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/repository"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/server"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/service"
	"github.com/spf13/cobra"
	"log"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "It establishes user service http connection",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("http called")

		cfg, err := config.GetConfig()
		if err != nil {
			log.Fatalf("error loading config: %v", err)
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

		authMiddleware := middlewares.NewAuthService(cfg)

		healthHandler := handlers.NewHealthCheckHandler()
		healthRoutes := routes.NewHealthRoutes(healthHandler)

		userRepository := repository.NewUserRepository(gormDB, gormDB, &log)

		userService := service.NewUserService(userRepository, nil)
		userHandler := handlers.NewUserHandler(userService, authMiddleware)
		userRoutes := routes.NewUserRoutes(userHandler)

		registerRoutes := routes.NewRegisterRoutes(
			routes.WithHealthRoute(healthRoutes),
			routes.WithUserRoute(userRoutes),
			routes.WithConfig(cfg),
		)

		httpServer := server.NewServer(
			server.WithHost(cfg.Server.Host),
			server.WithPort(cfg.Server.Port),
			server.WithApp(registerRoutes.RegisterRoutes()),
			server.WithLogger(&log),
		)

		log.Info().Str("port", cfg.Server.Port).Msg("starting http server")
		if err := httpServer.Connect(); err != nil {
			log.Fatal().Err(err).Msg("failed to start http server")
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
