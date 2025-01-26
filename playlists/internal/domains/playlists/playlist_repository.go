package playlists

import (
	"github.com/google/uuid"

	"github.com/ardwiinoo/micro-music/playlists/internal/domains/playlists/entities"
)

type PlaylistRepository interface {
	AddPlaylist(playlist *entities.AddPlaylist) (playlistId uuid.UUID, err error)
	GetPlaylistByUserPublicID(publicID uuid.UUID) (playlists []entities.DetailPlaylist, err error)
	DeletePlaylistByPlaylistIDAndUserPublicID(playlistID uuid.UUID, publicID uuid.UUID) (err error)
}