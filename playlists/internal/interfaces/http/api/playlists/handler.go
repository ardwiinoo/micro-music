package playlists

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/ardwiinoo/micro-music/playlists/internal/commons/exceptions"
	"github.com/ardwiinoo/micro-music/playlists/internal/domains/playlists/entities"
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

func (h *playlistHandler) AddPlaylistHandler(ctx *fiber.Ctx) error {
	var payload = entities.AddPlaylist{}

	if err := ctx.BodyParser(&payload); err != nil {
		return exceptions.InvariantError("invalid payload")
	}

	playlistID, err := h.container.AddPlaylistUseCase.Execute(ctx.UserContext(), payload)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"playlist_id": playlistID,
		},
	})
}

func (h *playlistHandler) DeletePlaylistHandler(ctx *fiber.Ctx) error {
	playlistID := ctx.Params("playlistId")

	playlistUUID, err := uuid.Parse(playlistID)
	if err != nil {
		return exceptions.InvariantError("invalid playlist ID")
	}

	err = h.container.DeletePlaylistUseCase.Execute(ctx.UserContext(), playlistUUID)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"status": "success",
	})
}

func (h *playlistHandler) AddPlaylistSongHandler(ctx *fiber.Ctx) error {
	playlistID := ctx.Params("playlistId")
	var payload = entities.AddPlaylistSong{}

	if err := ctx.BodyParser(&payload); err != nil {
		return exceptions.InvariantError("invalid payload")
	}

	playlistUUID, err := uuid.Parse(playlistID)
	if err != nil {
		return exceptions.InvariantError("invalid playlist ID")
	}

	err = h.container.AddPlaylistSongUseCase.Execute(ctx.UserContext(), playlistUUID, payload.SongID)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
	})
}
