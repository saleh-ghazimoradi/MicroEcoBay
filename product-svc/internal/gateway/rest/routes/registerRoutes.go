package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/gateway/rest/handlers"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/repository"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/service"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	v1 := app.Group("/v1")

	productRepository := repository.NewCatalogRepository(db)
	productService := service.NewCatalogService(productRepository)
	productHandler := handlers.NewCatalogHandler(productService)

	CatalogRoutes(v1, productHandler)
}
