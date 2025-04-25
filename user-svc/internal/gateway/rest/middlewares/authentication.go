package middlewares

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/config"
	"time"
)

func AuthMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		return ctx.Next()
	}
}

func GenerateToken(userId uint, email string) (string, error) {
	if userId == 0 || email == "" {
		return "", errors.New("userId and email are required")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"exp":     jwt.NewNumericDate(time.Now().Add(config.AppConfig.JWT.Exp)),
	})

	tokenString, err := token.SignedString([]byte(config.AppConfig.JWT.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
