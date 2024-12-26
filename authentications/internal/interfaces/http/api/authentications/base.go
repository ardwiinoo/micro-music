package authentications

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures"
)

func Init(router fiber.Router, container infrastructures.Container) {

	handler := newAuthenticationHandler(container)
	authenticationRouter := router.Group("auth")

	{
		authenticationRouter.Get("/login", handler.Hello)
	}
}