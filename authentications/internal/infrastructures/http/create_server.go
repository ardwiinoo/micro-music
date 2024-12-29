package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/micro-music/authentications/config"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures/http/middlewares"
	"github.com/ardwiinoo/micro-music/authentications/internal/interfaces/http/api/authentications"
	"github.com/ardwiinoo/micro-music/authentications/internal/interfaces/http/api/users"
)

func CreateServer(container *infrastructures.Container) *fiber.App {
	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
		ErrorHandler: middlewares.ErrorHandler,
	})

	router.Use(middlewares.Logger())

	users.Init(router, *container)
	authentications.Init(router, *container)

	return router
}