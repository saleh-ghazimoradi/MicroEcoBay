package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/gateway/rest/handlers"
)

func healthRoute(route fiber.Router, healthHandler *handlers.HealthHandler) {
	route.Get("/health", healthHandler.Health)
}
