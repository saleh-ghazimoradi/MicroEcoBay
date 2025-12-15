package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/gateway/rest/handlers"
)

type CatalogRoutes struct {
	catalogHandler *handlers.CatalogHandler
}

func (c *CatalogRoutes) CatalogRoute(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Post("/categories", c.catalogHandler.CreateCategory)
	v1.Get("/categories", c.catalogHandler.GetAllCategories)
	v1.Get("/categories/:id", c.catalogHandler.GetCategoryById)
	v1.Patch("/categories/:id", c.catalogHandler.UpdateCategory)

	v1.Post("/products", c.catalogHandler.CreateProduct)
	v1.Get("/products", c.catalogHandler.GetAllProducts)
	v1.Get("/products/:id", c.catalogHandler.GetProductById)
	v1.Patch("/products/:id", c.catalogHandler.UpdateProduct)

}

func NewCatalogRoues(catalogHandler *handlers.CatalogHandler) *CatalogRoutes {
	return &CatalogRoutes{
		catalogHandler: catalogHandler,
	}
}
