package helper

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type PaginatedResponse struct {
	Response Response      `json:"response"`
	Meta     PaginatedMeta `json:"meta"`
}

type PaginatedMeta struct {
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	Total     int64 `json:"total"`
	TotalPage int   `json:"total_page"`
}

func SuccessResponse(ctx *fiber.Ctx, message string, data any) error {
	return ctx.Status(fiber.StatusOK).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func CreatedResponse(ctx *fiber.Ctx, message string, data any) error {
	return ctx.Status(fiber.StatusCreated).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(ctx *fiber.Ctx, statusCode int, message string, err error) error {
	response := Response{
		Success: false,
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
	}

	return ctx.Status(statusCode).JSON(response)
}

func BadRequestResponse(ctx *fiber.Ctx, message string, err error) error {
	return ErrorResponse(ctx, fiber.StatusBadRequest, message, err)
}

func UnauthorizedResponse(ctx *fiber.Ctx, message string) error {
	return ErrorResponse(ctx, fiber.StatusUnauthorized, message, nil)
}

func ForbiddenResponse(ctx *fiber.Ctx, message string) error {
	return ErrorResponse(ctx, fiber.StatusForbidden, message, nil)
}

func NotFoundResponse(ctx *fiber.Ctx, message string) error {
	return ErrorResponse(ctx, fiber.StatusNotFound, message, nil)
}

func InternalServerError(ctx *fiber.Ctx, message string, err error) error {
	return ErrorResponse(ctx, fiber.StatusInternalServerError, message, err)
}

func PaginatedSuccessResponse(
	ctx *fiber.Ctx,
	message string,
	data any,
	meta PaginatedMeta,
) error {
	return ctx.Status(fiber.StatusOK).JSON(PaginatedResponse{
		Response: Response{
			Success: true,
			Message: message,
			Data:    data,
		},
		Meta: meta,
	})
}
