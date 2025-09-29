package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/dto"
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
		return 0, fiber.NewError(fiber.StatusBadRequest, "missing or invalid X-User-Id header")
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
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := c.cartService.Add(ctx.Context(), userId, payload); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "successfully created"})
}

func (c *CartHandler) GetCart(ctx *fiber.Ctx) error {
	userId, err := c.fetchAuthorizedUser(ctx)
	if err != nil {
		return err
	}

	cart, err := c.cartService.Get(ctx.Context(), userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": cart})
}

func (c *CartHandler) UpdateCart(ctx *fiber.Ctx) error {
	userId, err := c.fetchAuthorizedUser(ctx)
	if err != nil {
		return err
	}

	var payload *dto.CartUpdateQty
	if err := ctx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := c.cartService.UpdateQty(ctx.Context(), userId, payload.ProductID, payload.Qty); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully updated"})

}

func (c *CartHandler) DeleteCart(ctx *fiber.Ctx) error {
	userId, err := c.fetchAuthorizedUser(ctx)
	if err != nil {
		return err
	}

	var productInput = ctx.Params("productId")
	pid, _ := strconv.Atoi(productInput)
	if pid <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "invalid product id")
	}

	if err := c.cartService.Remove(ctx.Context(), userId, uint(pid)); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully deleted"})
}

func NewCartHandler(cartService service.CartService) *CartHandler {
	return &CartHandler{
		cartService: cartService,
	}
}
