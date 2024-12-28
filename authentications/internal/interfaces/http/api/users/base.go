package users

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures"
)

func Init(router fiber.Router, container infrastructures.Container) {

	handler := newUserHandler(container)
	userRouter := router.Group("users")
	
	{
		userRouter.Post("", handler.postUserHandler)
	}
}