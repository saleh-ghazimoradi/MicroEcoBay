package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/gateway/rest/handlers"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/gateway/rest/middlewares"
)

func userRoutes(v1 fiber.Router, handler *handlers.UserHandler) {
	v1.Group("/")

	v1.Post("/register", handler.Register)
	v1.Post("/login", handler.Login)
	v1.Post("/forgot-password", handler.ForgotPassword)
	v1.Post("/set-password", handler.SetPassword)

	v1.Use(middlewares.AuthMiddleware())

	v1.Post("/profile", handler.CreateProfile)
	v1.Get("/profile", handler.GetProfile)
	v1.Get("/auth", handler.Authentication)
	v1.Get("/me", handler.Me)
}
