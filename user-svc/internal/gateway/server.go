package gateway

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/config"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/infra/queue"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/gateway/grpc/order"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/slg"
)

func Server() error {

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Accept, Authorization",
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": "I'm breathing",
		})
	})

	kafkaProducer := queue.NewProducer(config.AppConfig.KafkaConfig.Broker, config.AppConfig.KafkaConfig.Topic)
	slg.Logger.Info("Kafka producer created", "kafka", kafkaProducer)

	orderServiceClient, err := order.NewGrpcOrderServiceClient("localhost:50051")
	if err != nil {
		slg.Logger.Error("Failed to create order service client", "error", err)
	}

	slg.Logger.Info("Order service client created", "client", orderServiceClient)

	if err := app.Listen(config.AppConfig.ServerConfig.Port); err != nil {
		return err
	}

	return nil
}
