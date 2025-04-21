package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/gateway/rest/handlers"
)

func userRoutes(v1 fiber.Router, handler *handlers.UserHandler) {
	v1.Group("/")
}
