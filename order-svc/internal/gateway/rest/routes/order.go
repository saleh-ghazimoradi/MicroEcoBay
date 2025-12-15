package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/gateway/rest/handlers"
)

type OrderRoutes struct {
	orderHandler *handlers.OrderHandler
}

func (o *OrderRoutes) OrderRoute(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Post("/orders", o.orderHandler.CreateOrder)
	v1.Get("/orders", o.orderHandler.GetOrderByUser)
	v1.Get("/orders/:orderId", o.orderHandler.GetOrderById)
}

func orderRoutes(route fiber.Router, orderHandler *handlers.OrderHandler) {

}

func NewOrderRoutes(orderHandler *handlers.OrderHandler) *OrderRoutes {
	return &OrderRoutes{
		orderHandler: orderHandler,
	}
}
