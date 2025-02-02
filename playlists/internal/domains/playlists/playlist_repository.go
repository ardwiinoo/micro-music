package playlists

import (
	"context"
	"github.com/google/uuid"

	"github.com/ardwiinoo/micro-music/playlists/internal/domains/playlists/entities"
)

type PlaylistRepository interface {
	AddPlaylist(ctx context.Context, playlist *entities.AddPlaylist, userID int) (playlistId uuid.UUID, err error)
	DeletePlaylist(ctx context.Context, playlistID uuid.UUID) (err error)
	ValidatePlaylistOwner(ctx context.Context, playlistID uuid.UUID, userID int) (err error)
	AddPlaylistSong(ctx context.Context, playlistID uuid.UUID, songID uuid.UUID) (err error)
	GetListPlaylistByUserPublicID(ctx context.Context, userID int) (playlists []entities.DetailPlaylist, err error)
}
