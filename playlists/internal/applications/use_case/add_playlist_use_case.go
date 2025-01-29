package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/ardwiinoo/micro-music/playlists/internal/commons/constants"
	"github.com/ardwiinoo/micro-music/playlists/internal/domains/playlists"
	"github.com/ardwiinoo/micro-music/playlists/internal/domains/playlists/entities"
	"github.com/ardwiinoo/micro-music/playlists/internal/domains/users"
)

type AddPlaylistUseCase interface {
	Execute(ctx context.Context, palyload entities.AddPlaylist) (playlistId uuid.UUID, err error)
}

type addPlaylistUseCase struct {
	playlistRepository playlists.PlaylistRepository
	userRepository users.UserRepository
}

func NewAddPlaylistUseCase(playlistRepository playlists.PlaylistRepository, userRepository users.UserRepository) AddPlaylistUseCase {
	return &addPlaylistUseCase{
		playlistRepository: playlistRepository,
		userRepository: userRepository,
	}
}

func (a *addPlaylistUseCase) Execute(ctx context.Context, payload entities.AddPlaylist) (playlistId uuid.UUID, err error) {
	publicID, ok := ctx.Value(constants.PublicIDContextKey).(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("ADD_PLAYLIST_USE_CASE.PUBLIC_ID_NOT_FOUND")
	}

	user, err := a.userRepository.GetDetailUserByPublicId(ctx, publicID.String())
	if err != nil {
		return uuid.Nil, errors.New("ADD_PLAYLIST_USE_CASE.USER_NOT_FOUND")
	}

	if user.Role != constants.UserRoleID {
		return uuid.Nil, errors.New("ADD_PLAYLIST_USE_CASE.USER_NOT_AUTHORIZED")
	}

	if err := payload.Validate(); err != nil {
		return uuid.Nil, err
	}

	return a.playlistRepository.AddPlaylist(&payload, user.ID)
}
