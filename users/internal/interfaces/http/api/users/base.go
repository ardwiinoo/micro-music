package users

import (
	"github.com/ardwiinoo/micro-music/users/internal/infrastructures"
	"github.com/ardwiinoo/micro-music/users/internal/infrastructures/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Init(router fiber.Router, container infrastructures.Container) {
	handler := NewUserHandler(container)

	userRouter := router.Group("/users", middlewares.TokenFilter(container.TokenManager))

	{
		userRouter.Post("/", handler.AddUserHandler)
		userRouter.Get("/", handler.GetListUserHandler)
		userRouter.Delete("/:id", handler.DeleteUserHandler)
	}
}
