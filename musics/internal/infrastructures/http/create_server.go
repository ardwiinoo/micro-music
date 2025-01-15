package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/micro-music/musics/config"
	"github.com/ardwiinoo/micro-music/musics/internal/infrastructures"
	"github.com/ardwiinoo/micro-music/musics/internal/infrastructures/http/middlewares"
)

func CreateServer(container *infrastructures.Container) *fiber.App {
	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
		ErrorHandler: middlewares.ErrorHandler,
	})

	router.Use(middlewares.Logger())
	router.Use(middlewares.AuthFilter(container.TokenManager))

	return router
}