package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/config"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/infra/queue"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/gateway/rest/handlers"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/gateway/rest/middlewares"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/repository"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/service"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/slg"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	v1 := app.Group("/v1")

	/* ---------- Dependencies ---------- */
	kafkaProducer := queue.NewProducer(config.AppConfig.KafkaConfig.Broker, config.AppConfig.KafkaConfig.Topic)
	slg.Logger.Info("Kafka producer created", "kafka", kafkaProducer)

	/* ---------- Repositories ---------- */
	userRepository := repository.NewUserRepository(db, db)

	/* ---------- Services ---------- */
	authService := middlewares.NewAuthService(config.AppConfig.JWT.Secret, config.AppConfig.JWT.Exp)
	userService := service.NewUserService(userRepository, kafkaProducer)

	/* ---------- Handlers ---------- */
	healthCheckHandler := handlers.NewHealthCheckHandler()
	userHandler := handlers.NewUserHandler(userService, authService)

	healthCheckRoute(v1, healthCheckHandler)
	userRoutes(v1, userHandler, authService)
}
