package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/ardwiinoo/micro-music/playlists/internal/commons/constants"
	"github.com/ardwiinoo/micro-music/playlists/internal/domains/playlists"
	"github.com/ardwiinoo/micro-music/playlists/internal/domains/users"
)

type DeletePlaylistUseCase interface {
	Execute(ctx context.Context, playlistID uuid.UUID) (err error)
}

type deletePlaylistUseCase struct {
	playlistRepository playlists.PlaylistRepository
	userRepository     users.UserRepository
}

// Execute implements DeletePlaylistUseCase.
func (d *deletePlaylistUseCase) Execute(ctx context.Context, playlistID uuid.UUID) (err error) {
	publicID, ok := ctx.Value(constants.PublicIDContextKey).(uuid.UUID)
	if !ok {
		return errors.New("DELETE_PLAYLIST_USE_CASE.PUBLIC_ID_NOT_FOUND")
	}

	user, err := d.userRepository.GetDetailUserByPublicId(ctx, publicID.String())
	if err != nil {
		return errors.New("DELETE_PLAYLIST_USE_CASE.USER_NOT_FOUND")
	}

	if user.Role != constants.UserRoleID {
		return errors.New("DELETE_PLAYLIST_USE_CASE.USER_NOT_AUTHORIZED")
	}

	if err := d.playlistRepository.ValidatePlaylistOwner(ctx, playlistID, user.ID); err != nil {
		return errors.New("DELETE_PLAYLIST_USE_CASE.PLAYLIST_NOT_FOUND")
	}

	return d.playlistRepository.DeletePlaylist(ctx, playlistID)
}

func NewDeletePlaylistUseCase(playlistRepository playlists.PlaylistRepository, userRepository users.UserRepository) DeletePlaylistUseCase {
	return &deletePlaylistUseCase{
		playlistRepository: playlistRepository,
		userRepository:     userRepository,
	}
}
