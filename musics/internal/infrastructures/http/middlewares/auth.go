package middlewares

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/micro-music/musics/internal/applications/security"
)

type contextKey string

const userContextKey contextKey = "user"

func AuthFilter(tokenManager security.TokenManager) fiber.Handler {
	return func(c *fiber.Ctx) error {
        token := c.Get("Authorization")
        if token == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid token"})
        }

        payload, err := tokenManager.VerifyToken(token)
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
        }

        ctx := context.WithValue(c.UserContext(), userContextKey, payload)
        c.SetUserContext(ctx)

        return c.Next()
    }
}