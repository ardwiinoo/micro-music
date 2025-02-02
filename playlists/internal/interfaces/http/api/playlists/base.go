package playlists

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/micro-music/playlists/internal/infrastructures"
	"github.com/ardwiinoo/micro-music/playlists/internal/infrastructures/http/middlewares"
)

func Init(router fiber.Router, container infrastructures.Container) {
	handler := NewPlaylistHandler(container)
	playlistRouter := router.Group("/playlists", middlewares.TokenFilter(container.TokenManager))

	{
		playlistRouter.Post("/", handler.AddPlaylistHandler)
		playlistRouter.Delete("/:playlistId", handler.DeletePlaylistHandler)
		playlistRouter.Post("/:playlistId/songs", handler.AddPlaylistSongHandler)
		playlistRouter.Get("/", handler.GetListPlaylistHandler)
	}
}
