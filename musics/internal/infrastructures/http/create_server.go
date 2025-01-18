package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	"github.com/ardwiinoo/micro-music/musics/config"
	"github.com/ardwiinoo/micro-music/musics/internal/infrastructures"
	"github.com/ardwiinoo/micro-music/musics/internal/infrastructures/http/middlewares"
	"github.com/ardwiinoo/micro-music/musics/internal/interfaces/songs"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9002
// @BasePath /
func CreateServer(container *infrastructures.Container) *fiber.App {
	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
		ErrorHandler: middlewares.ErrorHandler,
	})

	router.Use(middlewares.Logger())
	router.Get("/swagger/*", swagger.HandlerDefault)
	
	songs.Init(router, *container)

	return router
}