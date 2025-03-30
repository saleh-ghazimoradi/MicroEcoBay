package gateway

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/config"
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

	//kafkaProducer := queue.NewProducer(config.AppConfig.KafkaConfig.Broker, config.AppConfig.KafkaConfig.Topic)
	//slg.Logger.Info("Kafka producer created", "kafka", kafkaProducer)

	if err := app.Listen(config.AppConfig.ServerConfig.Port); err != nil {
		return err
	}

	return nil
}
