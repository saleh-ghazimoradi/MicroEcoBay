package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/service"
)

type CatalogHandler struct {
	catalogService service.CatalogService
}

func (c *CatalogHandler) CreateCategory(ctx *fiber.Ctx) error {
	return nil
}

func (c *CatalogHandler) GetAllCategories(ctx *fiber.Ctx) error {
	return nil
}

func (c *CatalogHandler) GetCategoryById(ctx *fiber.Ctx) error {
	return nil
}

func (c *CatalogHandler) UpdateCategory(ctx *fiber.Ctx) error {
	return nil
}

func (c *CatalogHandler) CreateProduct(ctx *fiber.Ctx) error {
	return nil
}

func (c *CatalogHandler) GetAllProducts(ctx *fiber.Ctx) error {
	return nil
}

func (c *CatalogHandler) GetProductById(ctx *fiber.Ctx) error {
	return nil
}

func (c *CatalogHandler) UpdateProduct(ctx *fiber.Ctx) error {
	return nil
}

func NewCatalogHandler(catalogService service.CatalogService) *CatalogHandler {
	return &CatalogHandler{
		catalogService: catalogService,
	}
}
