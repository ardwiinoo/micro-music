package authentications

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures"
)

func Init(router fiber.Router, container infrastructures.Container) {

	handler := newAuthenticationHandler(container)
	authenticationRouter := router.Group("auth")

	{
		authenticationRouter.Post("/login", handler.LoginHandler)
		authenticationRouter.Post("/register", handler.RegisterHandler)
		authenticationRouter.Get("/activate/:publicId", handler.ActivateUserHandler)
	}
}