package middlewares

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/ardwiinoo/micro-music/musics/internal/applications/security"
	"github.com/ardwiinoo/micro-music/musics/internal/commons/constants"
	"github.com/ardwiinoo/micro-music/musics/internal/commons/exceptions"

)

func AuthFilter(tokenManager security.TokenManager) fiber.Handler {
    return func(c *fiber.Ctx) error {
        token := c.Get("Authorization")
        if token == "" {
            return exceptions.UnauthorizedError("Token is required")
        }

        payload, err := tokenManager.VerifyToken(token)
        if (err != nil) {
            return exceptions.UnauthorizedError("Invalid token")
        }

        publicIDStr, ok := payload["public_id"].(string)
        if !ok {
            return exceptions.UnauthorizedError("Invalid token")
        }

        publicID, err := uuid.Parse(publicIDStr)
        if err != nil {
            return exceptions.UnauthorizedError("Invalid token")
        }

        ctx := context.WithValue(c.UserContext(), constants.PublicIDContextKey, publicID)
        c.SetUserContext(ctx)

        return c.Next()
    }
}