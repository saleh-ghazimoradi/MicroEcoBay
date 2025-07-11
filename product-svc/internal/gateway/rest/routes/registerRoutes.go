package routes

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/config"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/infra/queue"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/gateway/grpc/product"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/gateway/rest/handlers"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/repository"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/service"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/slg"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	v1 := app.Group("/v1")

	healthCheckHandler := handlers.NewHealthCheckHandler()
	catalogRepository := repository.NewCatalogRepository(db, db)
	catalogService := service.NewCatalogService(catalogRepository)
	catalogHandler := handlers.NewCatalogHandler(catalogService)

	kafkaConsumer := queue.NewConsumer(catalogService, config.AppConfig.KafkaConfig.Broker, "product-update-topic", "catalog-service-group")
	go kafkaConsumer.Listen(context.Background())

	go func() {
		list, err := net.Listen("tcp", ":50052")
		if err != nil {
			slg.Logger.Error("failed to listen grpc", "error", err.Error())
		}
		grpcServer := grpc.NewServer()
		productServer := product.NewProductGRPCService(catalogService)
		product.RegisterProductServiceServer(grpcServer, productServer)
		if err := grpcServer.Serve(list); err != nil {
			slg.Logger.Error("failed to start grpc server", "error", err.Error())
		}
	}()

	healthCheckRoute(v1, healthCheckHandler)
	CatalogRoutes(v1, catalogHandler)
}
