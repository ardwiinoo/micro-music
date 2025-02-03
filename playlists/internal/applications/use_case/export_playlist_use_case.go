package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ardwiinoo/micro-music/playlists/internal/commons/constants"
	"github.com/ardwiinoo/micro-music/playlists/internal/domains/playlists"
	"github.com/ardwiinoo/micro-music/playlists/internal/domains/users"
	"github.com/ardwiinoo/micro-music/playlists/internal/infrastructures/broker/rabbitmq"
	"github.com/google/uuid"
)

type ExportPlaylistUseCase interface {
	Execute(ctx context.Context, playlistID uuid.UUID) (err error)
}

type exportPlaylistUseCase struct {
	playlistRepository playlists.PlaylistRepository
	userRepository     users.UserRepository
	rabbitMQ           *rabbitmq.RabbitMQ
}

func NewExportPlaylistUseCase(playlistRepository playlists.PlaylistRepository, userRepository users.UserRepository, rabbitMQ *rabbitmq.RabbitMQ) ExportPlaylistUseCase {
	return &exportPlaylistUseCase{
		playlistRepository: playlistRepository,
		userRepository:     userRepository,
		rabbitMQ:           rabbitMQ,
	}
}
func (e exportPlaylistUseCase) Execute(ctx context.Context, playlistID uuid.UUID) (err error) {
	publicID, ok := ctx.Value(constants.PublicIDContextKey).(uuid.UUID)
	if !ok {
		return errors.New("EXPORT_PLAYLIST_USE_CASE.PUBLIC_ID_NOT_FOUND")
	}

	user, err := e.userRepository.GetDetailUserByPublicId(ctx, publicID.String())
	if err != nil {
		return errors.New("EXPORT_PLAYLIST_USE_CASE.USER_NOT_FOUND")
	}

	playlist, err := e.playlistRepository.GetPlaylistWithSongs(ctx, playlistID, user.ID)
	if err != nil {
		return err
	}

	playlistJSON, err := json.Marshal(playlist)
	if err != nil {
		return err
	}

	eventPayload := map[string]string{
		"email":       user.Email,
		"playlist_id": playlist.PlaylistID.String(),
		"playlist":    string(playlistJSON),
	}

	err = e.rabbitMQ.PublishEvent(constants.QueueTypes.EmailQueue, constants.EventType.ExportPlaylist, eventPayload)
	if err != nil {
		return err
	}

	return nil
}
