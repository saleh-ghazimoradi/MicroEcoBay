package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/gateway/rest/handlers"
)

func healthCheckRoute(v1 fiber.Router, handler *handlers.HealthCheckHandler) {
	v1.Get("/health", handler.HealthCheck)
}
