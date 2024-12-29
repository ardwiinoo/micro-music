package users

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/micro-music/authentications/internal/commons/exceptions"
	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users/entities"
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

func (h *userHandler) postUserHandler(ctx *fiber.Ctx) error {
	var payload = entities.RegisterUser{}

	if err := ctx.BodyParser(&payload); err != nil {
		return exceptions.InvariantError("invalid payload")
	}

	userId, err := h.container.RegisterUserUseCase.Execute(ctx.UserContext(), payload)

	if err != nil {
		return err
	}
	
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"userId": userId,
		},
	})
}