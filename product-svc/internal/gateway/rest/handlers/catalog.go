package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/dto"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/gateway/validator"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/service"
	"strconv"
)

type CatalogHandler struct {
	catalogService service.CatalogService
}

func (c *CatalogHandler) CreateCategory(ctx *fiber.Ctx) error {
	var payload *dto.CreateCategoryReq
	if err := ctx.BodyParser(&payload); err != nil {
		return badRequestResponse(ctx, err)
	}

	if err := validator.Validate.Struct(payload); err != nil {
		return badRequestResponse(ctx, err)
	}

	if err := c.catalogService.CreateCategory(ctx.Context(), payload); err != nil {
		return serverErrorResponse(ctx, err)
	}
	return successResponse(ctx, fiber.StatusCreated, "category created", payload)
}

func (c *CatalogHandler) GetAllCategories(ctx *fiber.Ctx) error {
	categories, err := c.catalogService.GetAllCategories(ctx.Context())
	if err != nil {
		return notFoundResponse(ctx)
	}

	return successResponse(ctx, fiber.StatusOK, "category list", categories)
}

func (c *CatalogHandler) GetCategoryById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	unId := uint(id)

	category, err := c.catalogService.GetCategoryById(ctx.Context(), unId)
	if err != nil {
		return notFoundResponse(ctx)
	}

	return successResponse(ctx, fiber.StatusOK, "category", category)
}

func (c *CatalogHandler) UpdateCategory(ctx *fiber.Ctx) error {
	var payload *dto.UpdateCategoryReq
	if err := ctx.BodyParser(&payload); err != nil {
		return badRequestResponse(ctx, err)
	}

	id, _ := strconv.Atoi(ctx.Params("id"))
	unId := uint(id)

	if err := c.catalogService.UpdateCategory(ctx.Context(), unId, payload); err != nil {
		return serverErrorResponse(ctx, err)
	}

	return successResponse(ctx, fiber.StatusOK, "category updated", payload)
}

func (c *CatalogHandler) CreateProduct(ctx *fiber.Ctx) error {
	var payload *dto.CreateProductReq
	if err := ctx.BodyParser(&payload); err != nil {
		return badRequestResponse(ctx, err)
	}

	if err := validator.Validate.Struct(payload); err != nil {
		return badRequestResponse(ctx, err)
	}

	if err := c.catalogService.CreateProduct(ctx.Context(), payload); err != nil {
		return serverErrorResponse(ctx, err)
	}

	return successResponse(ctx, fiber.StatusCreated, "product created", payload)
}

func (c *CatalogHandler) GetAllProducts(ctx *fiber.Ctx) error {
	products, err := c.catalogService.GetAllProducts(ctx.Context())
	if err != nil {
		return notFoundResponse(ctx)
	}

	return successResponse(ctx, fiber.StatusOK, "product list", products)
}

func (c *CatalogHandler) GetProductById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	unId := uint(id)

	product, err := c.catalogService.GetProductById(ctx.Context(), unId)
	if err != nil {
		return notFoundResponse(ctx)
	}

	return successResponse(ctx, fiber.StatusOK, "product", product)
}

func (c *CatalogHandler) UpdateProduct(ctx *fiber.Ctx) error {
	var payload *dto.UpdateProductReq
	if err := ctx.BodyParser(&payload); err != nil {
		return badRequestResponse(ctx, err)
	}

	id, _ := strconv.Atoi(ctx.Params("id"))
	unId := uint(id)

	if err := c.catalogService.UpdateProduct(ctx.Context(), unId, payload); err != nil {
		return serverErrorResponse(ctx, err)
	}

	return successResponse(ctx, fiber.StatusOK, "product updated", payload)
}

func NewCatalogHandler(catalogService service.CatalogService) *CatalogHandler {
	return &CatalogHandler{
		catalogService: catalogService,
	}
}
