package authentications

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

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

// LoginHandler godoc
// @Summary      Login user
// @Description  Authenticate user and generate access token
// @Tags         Authentications
// @Accept       json
// @Produce      json
// @Param        request body entities.LoginUser true "Login Payload"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /auth/login [post]
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

// RegisterHandler godoc
// @Summary      Register user
// @Description  Register new user
// @Tags         Authentications
// @Accept       json
// @Produce      json
// @Param        request body entities.RegisterUser true "Register Payload"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /auth/register [post]
func (h *authenticationHandler) RegisterHandler(ctx *fiber.Ctx) error {
	var payload = entities.RegisterUser{}

	if err := ctx.BodyParser(&payload); err != nil {
		return exceptions.InvariantError("invalid payload")
	}

	publicId, err := h.container.RegisterUserUseCase.Execute(ctx.UserContext(), payload)

	if err != nil {
		return err
	}
	
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"publicId": publicId,
		},
	})
}

// ActivateUserHandler godoc
// @Summary      Activate user
// @Description  Activate user account
// @Tags         Authentications
// @Accept       json
// @Produce      json
// @Param        publicId path string true "Public ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /auth/activate/{publicId} [get]
func (h *authenticationHandler) ActivateUserHandler(ctx *fiber.Ctx) error {
	publicId := ctx.Params("publicId")

	if publicId == "" {
		return exceptions.InvariantError("publicId is required")
	}

	uuidPublicId, err := uuid.Parse(publicId)
	if err != nil {
		return exceptions.InvariantError("invalid publicId format")
	}

	err = h.container.ActivateUserUseCase.Execute(ctx.UserContext(), uuidPublicId)

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"message": "User activated successfully",
	})
}