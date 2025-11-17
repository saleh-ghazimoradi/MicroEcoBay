package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/gateway/rest/handlers"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/repository"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/service"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	v1 := app.Group("/v1")

	healthHandler := handlers.NewHealthHandler()

	cartRepository := repository.NewCartRepository(db, db)
	cartService := service.NewCartService(cartRepository)
	cartHandler := handlers.NewCartHandler(cartService)
	orderRepository := repository.NewOrderRepository(db, db)
	orderService := service.NewOrderService(orderRepository)
	orderHandler := handlers.NewOrderHandler(orderService)

	healthRoute(v1, healthHandler)
	cartRoutes(v1, cartHandler)
	orderRoutes(v1, orderHandler)
}
