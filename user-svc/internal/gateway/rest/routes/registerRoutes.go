package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Register struct {
	healthRoute *HealthRoutes
	userRoute   *UserRoutes
}

type Options func(*Register)

func WithHealthRoute(healthRoute *HealthRoutes) Options {
	return func(r *Register) {
		r.healthRoute = healthRoute
	}
}

func WithUserRoute(userRoute *UserRoutes) Options {
	return func(r *Register) {
		r.userRoute = userRoute
	}
}

func (r *Register) RegisterRoutes() *fiber.App {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE",
	}))

	r.healthRoute.HealthRoute(app)
	r.userRoute.UserRoute(app)

	return app
}

func NewRegisterRoutes(opts ...Options) *Register {
	r := &Register{}
	for _, o := range opts {
		o(r)
	}
	return r
}
