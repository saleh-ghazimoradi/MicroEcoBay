package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/dto"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/gateway/rest/middlewares"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/helper"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/service"
)

type UserHandler struct {
	userService    service.UserService
	authMiddleware *middlewares.AuthService
}

func (u *UserHandler) Register(ctx *fiber.Ctx) error {
	payload := dto.UserSignup{}
	if err := ctx.BodyParser(&payload); err != nil {
		return helper.BadRequestResponse(ctx, "given invalid payload", err)
	}

	if err := u.userService.Register(ctx.Context(), &payload); err != nil {
		return helper.InternalServerError(ctx, "failed to register a user", err)
	}

	return helper.SuccessResponse(ctx, "user successfully created", nil)
}

func (u *UserHandler) Login(ctx *fiber.Ctx) error {
	payload := dto.UserLogin{}
	if err := ctx.BodyParser(&payload); err != nil {
		return helper.BadRequestResponse(ctx, "given invalid payload", err)
	}

	user, err := u.userService.Login(ctx.Context(), &payload)
	if err != nil {
		return helper.InternalServerError(ctx, "failed to login", err)
	}

	token, err := u.authMiddleware.GenerateToken(user.ID, user.Email)
	if err != nil {
		return helper.InternalServerError(ctx, "failed to generate token", err)
	}

	return helper.SuccessResponse(ctx, "user successfully login", token)
}

func (u *UserHandler) ForgotPassword(ctx *fiber.Ctx) error {
	payload := dto.ForgotPassword{}
	if err := ctx.BodyParser(&payload); err != nil {
		return helper.BadRequestResponse(ctx, "given invalid payload", err)
	}

	if err := u.userService.ForgotPassword(ctx.Context(), &payload); err != nil {
		return helper.InternalServerError(ctx, "failed to send password reset link", err)
	}

	return helper.SuccessResponse(ctx, "password reset link sent", nil)
}

func (u *UserHandler) SetPassword(ctx *fiber.Ctx) error {
	payload := dto.SetPassword{}

	if err := ctx.BodyParser(&payload); err != nil {
		return helper.BadRequestResponse(ctx, "given invalid payload", err)
	}

	if err := u.userService.SetPassword(ctx.Context(), &payload); err != nil {
		return helper.InternalServerError(ctx, "failed to set password", err)
	}

	return helper.SuccessResponse(ctx, "password reset link sent", nil)
}

func (u *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	payload := dto.UserProfile{}
	if err := ctx.BodyParser(&payload); err != nil {
		return helper.BadRequestResponse(ctx, "given invalid payload", err)
	}

	userId := ctx.Locals("userId")
	if userId == nil {
		return helper.UnauthorizedResponse(ctx, "user ID not found in context")
	}
	payload.UserId = userId.(uint)

	if err := u.userService.CreateProfile(ctx.Context(), &payload); err != nil {
		return helper.InternalServerError(ctx, "failed to create profile", err)
	}

	return helper.CreatedResponse(ctx, "user successfully created", nil)
}

func (u *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)
	profile, err := u.userService.GetProfile(ctx.Context(), userId)

	if err != nil {
		return helper.NotFoundResponse(ctx, "profile does not exist")
	}

	return helper.SuccessResponse(ctx, "profile successfully retrieved", profile)
}

func (u *UserHandler) Authentication(ctx *fiber.Ctx) error {
	user, err := u.userService.Authenticate(ctx)
	if err != nil {
		return helper.UnauthorizedResponse(ctx, "unauthorized")
	}

	return helper.SuccessResponse(ctx, "user successfully authenticated", user)
}

func (u *UserHandler) Me(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")
	if userId == nil {
		return helper.UnauthorizedResponse(ctx, "unauthorized")
	}
	userIdVal := userId.(uint)

	user, err := u.userService.GetProfile(ctx.Context(), userIdVal)
	if err != nil {
		return helper.NotFoundResponse(ctx, "profile does not exist")
	}

	return helper.SuccessResponse(ctx, "user successfully retrieved", user)
}

func NewUserHandler(userService service.UserService, authMiddleware *middlewares.AuthService) *UserHandler {
	return &UserHandler{
		userService:    userService,
		authMiddleware: authMiddleware,
	}
}
