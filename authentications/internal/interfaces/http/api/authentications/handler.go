package authentications

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users/entities"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures"
)

type authenticationHandler struct {
	container infrastructures.Container
}

func newAuthenticationHandler(container infrastructures.Container) *authenticationHandler {
	return &authenticationHandler{
		container: container,
	}
}

func (a *authenticationHandler) Hello(ctx *fiber.Ctx) error {

	request := entities.LoginUser {
		Email:    "test@example.com",
		Password: "password123",
	}

	token, err := a.container.LoginUserUseCase.Execute(ctx.Context(), request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"message": "Hello World",
		"token": token,
	})
}