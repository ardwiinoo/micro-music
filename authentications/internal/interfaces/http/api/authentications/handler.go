package authentications

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/micro-music/authentications/internal/commons/exceptions"
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

func (a *authenticationHandler) LoginHandler(ctx *fiber.Ctx) error {
	var payload = entities.LoginUser{}

	if err := ctx.BodyParser(&payload); err != nil {
		return exceptions.InvariantError("invalid payload")
	}	
	
	token, err := a.container.LoginUserUseCase.Execute(ctx.Context(), payload)
	
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"token": token,
		},
	})    
}