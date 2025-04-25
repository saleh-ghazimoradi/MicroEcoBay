package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func (u *UserHandler) Register(ctx *fiber.Ctx) error {
	return nil
}

func (u *UserHandler) Login(ctx *fiber.Ctx) error {
	return nil
}

func (u *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return nil
}

func (u *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	return nil
}

func (u *UserHandler) ForgotPassword(ctx *fiber.Ctx) error {
	return nil
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
