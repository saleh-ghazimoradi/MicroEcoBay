package middlewares

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/config"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/internal/dto"
	"strings"
	"time"
)

func AuthMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		authHeader := ctx.Get("Authorization")

		user, err := verifyToken(authHeader)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"error": err,
			})
		}

		ctx.Locals("userId", user.UserId)
		ctx.Locals("user", user)

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

func verifyToken(tokenString string) (*dto.AuthResponse, error) {
	tokenArr := strings.Split(tokenString, " ")

	if len(tokenArr) != 2 || tokenArr[0] != "Bearer" {
		return nil, errors.New("invalid token")
	}

	tokenStr := tokenArr[1]

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.AppConfig.JWT.Secret), nil
	})

	if err != nil {
		return nil, errors.New("token parse error")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, errors.New("token expired")
	}

	return &dto.AuthResponse{
		UserId: uint(claims["user_id"].(float64)),
		Email:  claims["email"].(string),
		Exp:    claims["exp"].(float64),
	}, nil
}
