package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/micro-music/authentications/config"
	"github.com/ardwiinoo/micro-music/authentications/internal/commons/exceptions/translator"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures"
	"github.com/ardwiinoo/micro-music/authentications/internal/interfaces/http/api/authentications"
	"github.com/ardwiinoo/micro-music/authentications/internal/interfaces/http/api/users"
)

func CreateServer(container *infrastructures.Container) *fiber.App {
	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	// Middleware
	router.Use(func(c *fiber.Ctx) error {
		err := c.Next()

		if err != nil {
			if mappedError, found := translator.ErrorMapping[err.Error()]; found {
				return c.Status(mappedError.HttpCode).JSON(mappedError)
			}

			return err
		}

		return nil
	})

	users.Init(router, *container)
	authentications.Init(router, *container)

	return router
}