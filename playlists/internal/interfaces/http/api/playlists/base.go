package playlists

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/micro-music/playlists/internal/infrastructures"
	"github.com/ardwiinoo/micro-music/playlists/internal/infrastructures/http/middlewares"
)

func Init(router fiber.Router, container infrastructures.Container) {
	handler := NewPlaylistHandler(container)
	playlistRouter := router.Group("/playlists")

	{
		playlistRouter.Post("/", middlewares.TokenFilter(container.TokenManager), handler.AddPlaylistHandler)
		playlistRouter.Delete("/:playlistId", middlewares.TokenFilter(container.TokenManager), handler.DeletePlaylistHandler)
		playlistRouter.Post("/:playlistId/songs", middlewares.TokenFilter(container.TokenManager), handler.AddPlaylistSongHandler)
	}
}
