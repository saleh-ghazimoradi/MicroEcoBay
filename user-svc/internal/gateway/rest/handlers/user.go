package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/dto"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/gateway/rest/middlewares"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/gateway/validator"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func (u *UserHandler) Register(ctx *fiber.Ctx) error {
	payload := dto.UserSignup{}
	if err := ctx.BodyParser(&payload); err != nil {
		return badRequestResponse(ctx, err)
	}

	if err := validator.Validate.Struct(payload); err != nil {
		return badRequestResponse(ctx, err)
	}

	if err := u.userService.Register(ctx.Context(), &payload); err != nil {
		return serverErrorResponse(ctx, err)
	}

	return successResponse(ctx, fiber.StatusCreated, "user successfully created", nil)
}

func (u *UserHandler) Login(ctx *fiber.Ctx) error {
	payload := dto.UserLogin{}
	if err := ctx.BodyParser(&payload); err != nil {
		return badRequestResponse(ctx, err)
	}

	if err := validator.Validate.Struct(payload); err != nil {
		return badRequestResponse(ctx, err)
	}

	user, err := u.userService.Login(ctx.Context(), &payload)
	if err != nil {
		return invalidCredentialsResponse(ctx)
	}

	token, err := middlewares.GenerateToken(user.ID, user.Email)
	if err != nil {
		return serverErrorResponse(ctx, err)
	}

	return successResponse(ctx, fiber.StatusOK, "user successfully login", token)
}

func (u *UserHandler) ForgotPassword(ctx *fiber.Ctx) error {
	payload := dto.ForgotPassword{}
	if err := ctx.BodyParser(&payload); err != nil {
		return badRequestResponse(ctx, err)
	}

	if err := validator.Validate.Struct(payload); err != nil {
		return badRequestResponse(ctx, err)
	}

	if err := u.userService.ForgotPassword(ctx.Context(), &payload); err != nil {
		return serverErrorResponse(ctx, err)
	}

	return successResponse(ctx, fiber.StatusOK, "password reset link sent", nil)
}

func (u *UserHandler) SetPassword(ctx *fiber.Ctx) error {
	payload := dto.SetPassword{}
	if err := ctx.BodyParser(&payload); err != nil {
		return badRequestResponse(ctx, err)
	}

	if err := validator.Validate.Struct(payload); err != nil {
		return badRequestResponse(ctx, err)
	}

	if err := u.userService.SetPassword(ctx.Context(), &payload); err != nil {
		return serverErrorResponse(ctx, err)
	}

	return successResponse(ctx, fiber.StatusOK, "password reset link sent", nil)
}

func (u *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	payload := dto.UserProfile{}
	if err := ctx.BodyParser(&payload); err != nil {
		return badRequestResponse(ctx, err)
	}

	userId := ctx.Locals("userId")
	payload.UserId = userId.(uint)

	if err := u.userService.CreateProfile(ctx.Context(), &payload); err != nil {
		return serverErrorResponse(ctx, err)
	}

	return successResponse(ctx, fiber.StatusCreated, "profile created successfully", nil)
}

func (u *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)
	profile, err := u.userService.GetProfile(ctx.Context(), userId)

	if err != nil {
		return serverErrorResponse(ctx, err)
	}

	return successResponse(ctx, fiber.StatusOK, "profile retrieved successfully", profile)
}

func (u *UserHandler) Authentication(ctx *fiber.Ctx) error {
	user, err := u.userService.Authenticate(ctx)
	if err != nil {
		return errorResponse(ctx, fiber.StatusUnauthorized, "unauthorized")
	}
	return successResponse(ctx, fiber.StatusOK, "user successfully authenticated", &fiber.Map{
		"authenticated": true,
		"user":          user,
	})
}

func (u *UserHandler) Me(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)

	user, err := u.userService.GetProfile(ctx.Context(), userId)
	if err != nil {
		return serverErrorResponse(ctx, err)
	}

	return successResponse(ctx, fiber.StatusOK, "user successfully authenticated", user)
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
