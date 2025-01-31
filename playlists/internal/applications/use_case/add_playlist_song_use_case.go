package usecase

import (
	"context"
	"errors"
	"github.com/ardwiinoo/micro-music/playlists/internal/commons/constants"
	"github.com/ardwiinoo/micro-music/playlists/internal/domains/playlists"
	"github.com/ardwiinoo/micro-music/playlists/internal/domains/users"
	"github.com/google/uuid"
)

type AddPlaylistSongUseCase interface {
	Execute(ctx context.Context, playlistID uuid.UUID, songID uuid.UUID) (err error)
}

type addPlaylistSongUseCase struct {
	playlistRepository playlists.PlaylistRepository
	userRepository     users.UserRepository
}

func (a addPlaylistSongUseCase) Execute(ctx context.Context, playlistID uuid.UUID, songID uuid.UUID) (err error) {
	publicID, ok := ctx.Value(constants.PublicIDContextKey).(uuid.UUID)
	if !ok {
		return errors.New("ADD_PLAYLIST_SONG_USE_CASE.PUBLIC_ID_NOT_FOUND")
	}

	user, err := a.userRepository.GetDetailUserByPublicId(ctx, publicID.String())
	if err != nil {
		return errors.New("ADD_PLAYLIST_SONG_USE_CASE.USER_NOT_FOUND")
	}

	if user.Role != constants.UserRoleID {
		return errors.New("ADD_PLAYLIST_SONG_USE_CASE.USER_NOT_AUTHORIZED")
	}

	if err := a.playlistRepository.ValidatePlaylistOwner(ctx, playlistID, user.ID); err != nil {
		return errors.New("ADD_PLAYLIST_SONG_USE_CASE.PLAYLIST_NOT_FOUND")
	}

	return a.playlistRepository.AddPlaylistSong(ctx, playlistID, songID)
}

func NewAddPlaylistSongUseCase(playlistRepository playlists.PlaylistRepository, userRepository users.UserRepository) AddPlaylistSongUseCase {
	return &addPlaylistSongUseCase{
		playlistRepository: playlistRepository,
		userRepository:     userRepository,
	}
}
