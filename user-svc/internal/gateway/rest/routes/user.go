package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/gateway/rest/handlers"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/gateway/rest/middlewares"
)

type UserRoutes struct {
	userHandler    *handlers.UserHandler
	authMiddleware *middlewares.AuthService
}

func (u *UserRoutes) UserRoute(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Post("/register", u.userHandler.Register)
	v1.Post("/login", u.userHandler.Login)
	v1.Post("/forgot-password", u.userHandler.ForgotPassword)
	v1.Post("/set-password", u.userHandler.SetPassword)

	v1.Use(u.authMiddleware.AuthMiddleware())
	v1.Post("/profile", u.userHandler.ForgotPassword)
	v1.Get("/profile", u.userHandler.GetProfile)
	v1.Get("/auth", u.userHandler.Authentication)
	v1.Get("/me", u.userHandler.Me)
}

func NewUserRoutes(userHandler *handlers.UserHandler) *UserRoutes {
	return &UserRoutes{
		userHandler: userHandler,
	}
}
