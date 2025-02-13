package songs

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/micro-music/musics/internal/infrastructures"
	"github.com/ardwiinoo/micro-music/musics/internal/infrastructures/http/middlewares"
)

func Init(router fiber.Router, container infrastructures.Container) {

	handler := NewSongHandler(container)
	songRouter := router.Group("/songs")

	{
		songRouter.Get("/", handler.getListSongHandler)
		songRouter.Post("/", middlewares.TokenFilter(container.TokenManager), handler.addSongHandler)
		songRouter.Get("/:id/stream", handler.StreamSongHandler)

	}
}
