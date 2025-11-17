package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/dto"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/service"
	"strconv"
)

type OrderHandler struct {
	orderService service.OrderService
}

func (o *OrderHandler) fetchAuthorizedUser(ctx *fiber.Ctx) (uint, error) {
	rawUserData := ctx.Get("X-User-Id")
	userId, err := strconv.Atoi(rawUserData)
	if err != nil || userId <= 0 {
		return 0, fiber.NewError(fiber.StatusBadRequest, "missing or invalid X-User-Id header")
	}
	return uint(userId), nil
}

func (o *OrderHandler) CreateOrder(ctx *fiber.Ctx) error {
	userId, err := o.fetchAuthorizedUser(ctx)
	if err != nil {
		return err
	}

	var input dto.CreateOrderRequest
	if err := ctx.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	input.UserId = userId

	order, err := o.orderService.CreateOrder(ctx.Context(), &input)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"order": order,
	})
}

func (o *OrderHandler) GetOrderById(ctx *fiber.Ctx) error {
	userId, err := o.fetchAuthorizedUser(ctx)
	if err != nil {
		return err
	}

	orderId, err := strconv.Atoi(ctx.Params("orderId"))
	if err != nil || orderId <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "missing or invalid orderId header")
	}

	order, err := o.orderService.GetOrderById(ctx.Context(), userId, uint(orderId))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"order": order,
	})
}

func (o *OrderHandler) GetOrderByUser(ctx *fiber.Ctx) error {
	userId, err := o.fetchAuthorizedUser(ctx)
	if err != nil {
		return err
	}

	orders, err := o.orderService.GetOrderByUser(ctx.Context(), userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"orders": orders,
	})
}

func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}
