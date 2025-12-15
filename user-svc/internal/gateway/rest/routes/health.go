package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/gateway/rest/handlers"
)

type HealthRoutes struct {
	healthHandler *handlers.HealthCheckHandler
}

func (h *HealthRoutes) HealthRoute(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Get("/health", h.healthHandler.HealthCheck)

}

func NewHealthRoutes(healthHandler *handlers.HealthCheckHandler) *HealthRoutes {
	return &HealthRoutes{
		healthHandler: healthHandler,
	}
}
