package usecase

import (
	"context"
	"errors"
	"github.com/ardwiinoo/micro-music/playlists/internal/commons/constants"
	"github.com/ardwiinoo/micro-music/playlists/internal/domains/playlists"
	"github.com/ardwiinoo/micro-music/playlists/internal/domains/playlists/entities"
	"github.com/ardwiinoo/micro-music/playlists/internal/domains/users"
	"github.com/google/uuid"
)

type GetListPlaylistUseCase interface {
	Execute(ctx context.Context) ([]entities.DetailPlaylist, error)
}

type getListPlaylistUseCase struct {
	playlistRepository playlists.PlaylistRepository
	userRepository     users.UserRepository
}

func NewGetListPlaylistUseCase(playlistRepository playlists.PlaylistRepository, userRepository users.UserRepository) GetListPlaylistUseCase {
	return &getListPlaylistUseCase{
		playlistRepository: playlistRepository,
		userRepository:     userRepository,
	}
}

func (g getListPlaylistUseCase) Execute(ctx context.Context) ([]entities.DetailPlaylist, error) {
	publicID, ok := ctx.Value(constants.PublicIDContextKey).(uuid.UUID)
	if !ok {
		return nil, errors.New("GET_LIST_PLAYLIST_USE_CASE.PUBLIC_ID_NOT_FOUND")
	}

	user, err := g.userRepository.GetDetailUserByPublicId(ctx, publicID.String())
	if err != nil {
		return nil, errors.New("GET_LIST_PLAYLIST_USE_CASE.USER_NOT_FOUND")
	}

	return g.playlistRepository.GetListPlaylistByUserPublicID(ctx, user.ID)
}
