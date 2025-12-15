package cmd

import (
	"fmt"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/config"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/infra/postgresql"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/gateway/rest/handlers"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/gateway/rest/routes"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/logger"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/repository"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/server"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/service"
	"log"

	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "It establishes order service http connection",
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

		cartRepository := repository.NewCartRepository(gormDB, gormDB)
		orderRepository := repository.NewOrderRepository(gormDB, gormDB)

		cartService := service.NewCartService(cartRepository)
		orderService := service.NewOrderService(orderRepository)

		healthHandler := handlers.NewHealthCheckHandler()
		cartHandler := handlers.NewCartHandler(cartService)
		orderHandler := handlers.NewOrderHandler(orderService)

		healthRoutes := routes.NewHealthRoutes(healthHandler)
		cartRoutes := routes.NewCartRoutes(cartHandler)
		orderRoutes := routes.NewOrderRoutes(orderHandler)

		registerRoutes := routes.NewRegisterRoutes(
			routes.WithHealthRoute(healthRoutes),
			routes.WithCartRoute(cartRoutes),
			routes.WithOrderRoute(orderRoutes),
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
