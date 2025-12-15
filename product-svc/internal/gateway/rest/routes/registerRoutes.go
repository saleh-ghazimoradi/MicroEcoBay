package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/config"
)

type Register struct {
	config       *config.Config
	healthRoute  *HealthRoutes
	catalogRoute *CatalogRoutes
}

type Options func(*Register)

func WithHealthRoute(healthRoute *HealthRoutes) func(*Register) {
	return func(r *Register) {
		r.healthRoute = healthRoute
	}
}

func WithCatalogRoute(catalogRoute *CatalogRoutes) func(*Register) {
	return func(r *Register) {
		r.catalogRoute = catalogRoute
	}
}

func (r *Register) RegisterRoutes() *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  r.config.Server.ReadTimeout,
		WriteTimeout: r.config.Server.WriteTimeout,
		IdleTimeout:  r.config.Server.IdleTimeout,
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE",
	}))

	r.healthRoute.HealthRoute(app)
	r.catalogRoute.CatalogRoute(app)

	return app
}

func NewRegister(opts ...Options) *Register {
	r := &Register{}
	for _, f := range opts {
		f(r)
	}
	return r
}
