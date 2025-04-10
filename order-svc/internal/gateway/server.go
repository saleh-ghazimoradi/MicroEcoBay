package gateway

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/config"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/slg"
)

func Server() error {

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-type, Accept, Authorization",
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": "I'm breathing",
		})
	})

	slg.Logger.Info("the server is running", "port", config.AppConfig.ServerConfig)

	if err := app.Listen(config.AppConfig.ServerConfig.Port); err != nil {
		return err
	}

	return nil
}
