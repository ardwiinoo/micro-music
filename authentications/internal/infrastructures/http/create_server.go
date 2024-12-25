package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/micro-music/authentications/config"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures"
	"github.com/ardwiinoo/micro-music/authentications/internal/interfaces/http/api/users"
)

func CreateServer(container *infrastructures.Container) *fiber.App {
	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	// Middleware

	// Init modules
	users.Init(router, *container)

	// Error Handling

	return router
}