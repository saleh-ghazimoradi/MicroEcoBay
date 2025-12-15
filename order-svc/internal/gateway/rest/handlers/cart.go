package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/dto"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/helper"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/service"
	"strconv"
)

type CartHandler struct {
	cartService service.CartService
}

func (c *CartHandler) fetchAuthorizedUser(ctx *fiber.Ctx) (uint, error) {
	rawUserData := ctx.Get("X-User-Id")
	userId, err := strconv.Atoi(rawUserData)
	if err != nil || userId <= 0 {
		return 0, helper.BadRequestResponse(ctx, "missing or invalid X-User_Id header", err)
	}
	return uint(userId), nil
}

func (c *CartHandler) CreateCart(ctx *fiber.Ctx) error {
	userId, err := c.fetchAuthorizedUser(ctx)
	if err != nil {
		return err
	}

	var payload *dto.Cart
	if err := ctx.BodyParser(&payload); err != nil {
		return helper.BadRequestResponse(ctx, "given invalid payload", err)
	}

	if err := c.cartService.Add(ctx.Context(), userId, payload); err != nil {
		return helper.InternalServerError(ctx, "failed to add cart", err)
	}

	return helper.CreatedResponse(ctx, "cart successfully created", nil)
}

func (c *CartHandler) GetCart(ctx *fiber.Ctx) error {
	userId, err := c.fetchAuthorizedUser(ctx)
	if err != nil {
		return err
	}

	cart, err := c.cartService.Get(ctx.Context(), userId)
	if err != nil {
		return helper.InternalServerError(ctx, "failed to get cart", err)
	}

	return helper.SuccessResponse(ctx, "cart successfully retrieved", cart)
}

func (c *CartHandler) UpdateCart(ctx *fiber.Ctx) error {
	userId, err := c.fetchAuthorizedUser(ctx)
	if err != nil {
		return err
	}

	var payload *dto.CartUpdateQty
	if err := ctx.BodyParser(&payload); err != nil {
		return helper.BadRequestResponse(ctx, "given invalid payload", err)
	}

	if err := c.cartService.UpdateQty(ctx.Context(), userId, payload.ProductID, payload.Qty); err != nil {
		return helper.InternalServerError(ctx, "failed to update cart", err)
	}

	return helper.SuccessResponse(ctx, "cart successfully updated", nil)
}

func (c *CartHandler) DeleteCart(ctx *fiber.Ctx) error {
	userId, err := c.fetchAuthorizedUser(ctx)
	if err != nil {
		return err
	}

	var productInput = ctx.Params("productId")
	pid, _ := strconv.Atoi(productInput)
	if pid <= 0 {
		return helper.BadRequestResponse(ctx, "invalid product id", err)
	}

	if err := c.cartService.Remove(ctx.Context(), userId, uint(pid)); err != nil {
		return helper.InternalServerError(ctx, "failed to delete cart", err)
	}

	return helper.SuccessResponse(ctx, "cart successfully deleted", nil)
}

func NewCartHandler(cartService service.CartService) *CartHandler {
	return &CartHandler{
		cartService: cartService,
	}
}
