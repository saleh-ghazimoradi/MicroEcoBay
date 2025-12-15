package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/gateway/rest/handlers"
)

type CartRoutes struct {
	cartHandler *handlers.CartHandler
}

func (c *CartRoutes) CarteRoute(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Post("/cart", c.cartHandler.CreateCart)
	v1.Get("/cart/:id", c.cartHandler.GetCart)
	v1.Patch("/cart", c.cartHandler.UpdateCart)
	v1.Delete("/cart/:productId", c.cartHandler.DeleteCart)
}

func NewCartRoutes(cartHandler *handlers.CartHandler) *CartRoutes {
	return &CartRoutes{
		cartHandler: cartHandler,
	}
}
