package handlers

import "github.com/gofiber/fiber/v2"

type HealthHandler struct{}

func (h *HealthHandler) Health(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "I'm breathing",
	})
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}
