package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/ardwiinoo/micro-music/playlists/config"
	"github.com/ardwiinoo/micro-music/playlists/internal/infrastructures"
	"github.com/ardwiinoo/micro-music/playlists/internal/infrastructures/http/middlewares"
)

func CreateServer(container *infrastructures.Container) *fiber.App {
	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
		ErrorHandler: middlewares.ErrorHandler,
	})

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	router.Use(logger.New(logger.Config{
    	Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	
	return router
}