package usecase

import (
	"context"

	"github.com/ardwiinoo/micro-music/playlists/internal/domains/playlists"
	"github.com/ardwiinoo/micro-music/playlists/internal/domains/playlists/entities"
	"github.com/google/uuid"
)

type AddPlaylistUseCase interface {
	Execute(ctx context.Context, palyload entities.AddPlaylist) (playlistId uuid.UUID, err error)
}

type addPlaylistUseCase struct {
	playlistRepository playlists.PlaylistRepository
}

func NewAddPlaylistUseCase(playlistRepository playlists.PlaylistRepository) AddPlaylistUseCase {
	return &addPlaylistUseCase{
		playlistRepository: playlistRepository,
	}
}

func (a *addPlaylistUseCase) Execute(ctx context.Context, payload entities.AddPlaylist) (playlistId uuid.UUID, err error) {
	playlistId, err = a.playlistRepository.AddPlaylist(&payload)
	if err != nil {
		return uuid.UUID{}, err
	}

	return playlistId, nil
}