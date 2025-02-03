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

// AddPlaylistHandler godoc
// @Summary      Add a new playlist
// @Description  Add a new playlist to the database
// @Tags         Playlists
// @Accept       json
// @Produce      json
// @Param        request body entities.AddPlaylist true "Playlist Payload"
// @Param        Authorization header string true "Authorization Bearer Token"
// @Success      201 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Security     ApiKeyAuth
// @Router       /playlists [post]
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

// DeletePlaylistHandler godoc
// @Summary      Delete a playlist
// @Description  Delete a playlist from the database
// @Tags         Playlists
// @Accept       json
// @Produce      json
// @Param        playlistId path string true "Playlist ID"
// @Param        Authorization header string true "Authorization Bearer Token"
// @Success      204 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Security     ApiKeyAuth
// @Router       /playlists/{playlistId} [delete]
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

// AddPlaylistSongHandler godoc
// @Summary      Add a song to a playlist
// @Description  Add a new song to an existing playlist
// @Tags         Playlists
// @Accept       json
// @Produce      json
// @Param        playlistId path string true "Playlist ID"
// @Param        request body entities.AddPlaylistSong true "Song Payload"
// @Param        Authorization header string true "Authorization Bearer Token"
// @Success      201 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Security     ApiKeyAuth
// @Router       /playlists/{playlistId}/songs [post]
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

// GetListPlaylistHandler godoc
// @Summary      Get list of playlists
// @Description  Retrieve a list of playlists available in the system
// @Tags         Playlists
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Authorization Bearer Token"
// @Success      200 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Security     ApiKeyAuth
// @Router       /playlists [get]
func (h *playlistHandler) GetListPlaylistHandler(ctx *fiber.Ctx) error {
	listPlaylist, err := h.container.GetListPlaylistUseCase.Execute(ctx.UserContext())
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   listPlaylist,
	})
}

// ExportPlaylistHandler godoc
// @Summary      Export playlist
// @Description  Export a playlist and send it via email
// @Tags         Playlists
// @Accept       json
// @Produce      json
// @Param        playlistId path string true "Playlist ID"
// @Param        Authorization header string true "Authorization Bearer Token"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Security     ApiKeyAuth
// @Router       /playlists/{playlistId}/export [post]
func (h *playlistHandler) ExportPlaylistHandler(ctx *fiber.Ctx) error {
	playlistID := ctx.Params("playlistId")

	playlistUUID, err := uuid.Parse(playlistID)
	if err != nil {
		return exceptions.InvariantError("invalid playlist ID")
	}

	err = h.container.ExportPlaylistUseCase.Execute(ctx.UserContext(), playlistUUID)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "playlist export initiated",
	})
}
