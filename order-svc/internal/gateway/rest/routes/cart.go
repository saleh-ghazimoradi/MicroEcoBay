package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/gateway/rest/handlers"
)

func cartRoutes(route fiber.Router, cartHandler *handlers.CartHandler) {
	route.Post("/cart", cartHandler.CreateCart)
	route.Get("/cart/:id", cartHandler.GetCart)
	route.Patch("/cart", cartHandler.UpdateCart)
	route.Delete("/cart/:productId", cartHandler.DeleteCart)
}
