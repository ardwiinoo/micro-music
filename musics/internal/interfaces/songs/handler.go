package songs

import (
	"github.com/gofiber/fiber/v2"
	"io"

	"github.com/ardwiinoo/micro-music/musics/internal/commons/exceptions"
	"github.com/ardwiinoo/micro-music/musics/internal/domains/songs/entities"
	"github.com/ardwiinoo/micro-music/musics/internal/infrastructures"
)

type songHandler struct {
	container infrastructures.Container
}

func NewSongHandler(container infrastructures.Container) *songHandler {
	return &songHandler{
		container: container,
	}
}

// GetListSongHandler godoc
// @Summary      Get list of songs
// @Description  Retrieve a list of songs available in the system
// @Tags         Songs
// @Accept       json
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /songs [get]
func (s *songHandler) getListSongHandler(ctx *fiber.Ctx) error {
	listSong, isCached, err := s.container.GetListSongUseCase.Execute(ctx.UserContext())
	if err != nil {
		return err
	}

	if isCached {
		ctx.Set("X-Cache-Status", "HIT")
	} else {
		ctx.Set("X-Cache-Status", "MISS")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   listSong,
	})
}

// AddSongHandler godoc
// @Summary      Add a new song
// @Description  Add a new song to the database
// @Tags         Songs
// @Accept       json
// @Produce      json
// @Param        request body entities.AddSong true "Song Payload"
// @Param        Authorization header string true "Authorization Bearer Token"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Security     ApiKeyAuth
// @Router       /songs [post]
func (s *songHandler) addSongHandler(ctx *fiber.Ctx) error {
	var payload = entities.AddSong{}

	if err := ctx.BodyParser(&payload); err != nil {
		return exceptions.InvariantError("invalid payload")
	}

	id, err := s.container.AddSongUseCase.Execute(ctx.UserContext(), payload)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"song_id": id,
		},
	})
}

func (s *songHandler) StreamSongHandler(ctx *fiber.Ctx) error {
	songID := ctx.Params("id")
	rangeHeader := ctx.Get("Range")

	body, statusCode, contentType, err := s.container.StreamSongUseCase.Execute(ctx.UserContext(), songID, rangeHeader)
	if err != nil {
		return err
	}

	defer body.Close()

	ctx.Set("Content-Type", contentType)
	ctx.Status(statusCode)

	_, err = io.Copy(ctx.Response().BodyWriter(), body)
	if err != nil {
		return err
	}
	
	return nil
}
