package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/gateway/rest/handlers"
)

func CatalogRoutes(routes fiber.Router, catalog *handlers.CatalogHandler) {
	routes.Post("/categories", catalog.CreateCategory)
	routes.Get("/categories", catalog.GetAllCategories)
	routes.Get("/categories/:id", catalog.GetCategoryById)
	routes.Patch("/categories/:id", catalog.UpdateCategory)

	routes.Post("/products", catalog.CreateProduct)
	routes.Get("/products", catalog.GetAllProducts)
	routes.Get("/products/:id", catalog.GetProductById)
	routes.Patch("/products/:id", catalog.UpdateProduct)
}
