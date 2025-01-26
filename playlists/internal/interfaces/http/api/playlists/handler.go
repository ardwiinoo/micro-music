package playlists

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/micro-music/playlists/internal/infrastructures"
)

type playlistHandler struct {
	container infrastructures.Container
}

func NewPlaylistHandler(container infrastructures.Container) *playlistHandler {
	return &playlistHandler{
		container: container,
	}
}

func (h *playlistHandler) AddPlaylistHandler(c *fiber.Ctx) error {
	
	return nil
}