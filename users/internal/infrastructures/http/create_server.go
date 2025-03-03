package http

import (
	"github.com/ardwiinoo/micro-music/users/config"
	"github.com/ardwiinoo/micro-music/users/internal/infrastructures"
	"github.com/ardwiinoo/micro-music/users/internal/infrastructures/http/middlewares"
	"github.com/ardwiinoo/micro-music/users/internal/interfaces/http/api/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func CreateServer(container *infrastructures.Container) *fiber.App {
	router := fiber.New(fiber.Config{
		Prefork:      true,
		AppName:      config.Cfg.App.Name,
		ErrorHandler: middlewares.ErrorHandler,
	})

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	router.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	users.Init(router, *container)

	return router
}
