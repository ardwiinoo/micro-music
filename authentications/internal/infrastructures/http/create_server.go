package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	"github.com/ardwiinoo/micro-music/authentications/config"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures/http/middlewares"
	authentications "github.com/ardwiinoo/micro-music/authentications/internal/interfaces/http/api/auth"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9001
// @BasePath /
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

	router.Get("/auth/swagger/*", swagger.HandlerDefault)

	authentications.Init(router, *container)

	return router
}