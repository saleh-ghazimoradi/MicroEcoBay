package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/dto"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/helper"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/service"
	"strconv"
)

type CatalogHandler struct {
	catalogService service.CatalogService
}

func (c *CatalogHandler) CreateCategory(ctx *fiber.Ctx) error {
	var payload *dto.CreateCategoryReq
	if err := ctx.BodyParser(&payload); err != nil {
		return helper.BadRequestResponse(ctx, "given invalid payload", err)
	}

	if err := c.catalogService.CreateCategory(ctx.Context(), payload); err != nil {
		return helper.InternalServerError(ctx, "failed to create category", err)
	}

	return helper.CreatedResponse(ctx, "category successfully created", payload)
}

func (c *CatalogHandler) GetAllCategories(ctx *fiber.Ctx) error {
	categories, err := c.catalogService.GetAllCategories(ctx.Context())
	if err != nil {
		return helper.NotFoundResponse(ctx, "Category not found")
	}

	return helper.SuccessResponse(ctx, "category successfully retrieved", categories)
}

func (c *CatalogHandler) GetCategoryById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	unId := uint(id)

	category, err := c.catalogService.GetCategoryById(ctx.Context(), unId)
	if err != nil {
		return helper.NotFoundResponse(ctx, "category not found")
	}

	return helper.SuccessResponse(ctx, "category successfully retrieved", category)
}

func (c *CatalogHandler) UpdateCategory(ctx *fiber.Ctx) error {
	var payload *dto.UpdateCategoryReq
	if err := ctx.BodyParser(&payload); err != nil {
		return helper.BadRequestResponse(ctx, "given invalid payload", err)
	}

	id, _ := strconv.Atoi(ctx.Params("id"))
	unId := uint(id)

	if err := c.catalogService.UpdateCategory(ctx.Context(), unId, payload); err != nil {
		return helper.InternalServerError(ctx, "failed to update category", err)
	}

	return helper.SuccessResponse(ctx, "category successfully updated", payload)
}

func (c *CatalogHandler) CreateProduct(ctx *fiber.Ctx) error {
	var payload *dto.CreateProductReq
	if err := ctx.BodyParser(&payload); err != nil {
		return helper.BadRequestResponse(ctx, "given invalid payload", err)
	}

	if err := c.catalogService.CreateProduct(ctx.Context(), payload); err != nil {
		return helper.InternalServerError(ctx, "failed to create product", err)
	}

	return helper.CreatedResponse(ctx, "product successfully created", payload)
}

func (c *CatalogHandler) GetAllProducts(ctx *fiber.Ctx) error {
	products, err := c.catalogService.GetAllProducts(ctx.Context())
	if err != nil {
		return helper.NotFoundResponse(ctx, "Product not found")
	}

	return helper.SuccessResponse(ctx, "product successfully retrieved", products)
}

func (c *CatalogHandler) GetProductById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	unId := uint(id)

	product, err := c.catalogService.GetProductById(ctx.Context(), unId)
	if err != nil {
		return helper.NotFoundResponse(ctx, "Product not found")
	}

	return helper.SuccessResponse(ctx, "product successfully retrieved", product)
}

func (c *CatalogHandler) UpdateProduct(ctx *fiber.Ctx) error {
	var payload *dto.UpdateProductReq
	if err := ctx.BodyParser(&payload); err != nil {
		return helper.BadRequestResponse(ctx, "given invalid payload", err)
	}

	id, _ := strconv.Atoi(ctx.Params("id"))
	unId := uint(id)

	if err := c.catalogService.UpdateProduct(ctx.Context(), unId, payload); err != nil {
		return helper.InternalServerError(ctx, "failed to update category", err)
	}

	return helper.SuccessResponse(ctx, "product successfully updated", payload)
}

func NewCatalogHandler(catalogService service.CatalogService) *CatalogHandler {
	return &CatalogHandler{
		catalogService: catalogService,
	}
}
