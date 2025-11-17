package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/gateway/rest/handlers"
)

func orderRoutes(route fiber.Router, orderHandler *handlers.OrderHandler) {
	route.Post("/orders", orderHandler.CreateOrder)
	route.Get("/orders", orderHandler.GetOrderByUser)
	route.Get("/orders/:orderId", orderHandler.GetOrderById)
}
