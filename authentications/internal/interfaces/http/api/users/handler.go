package users

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures"
)

type userHandler struct {
	container infrastructures.Container
}

func newUserHandler(container infrastructures.Container) *userHandler {
	return &userHandler{
		container: container,
	}
}

func (h *userHandler) Hello(ctx *fiber.Ctx) error {
	
	return ctx.JSON(fiber.Map{
		"message": "Hello World",
	})
}