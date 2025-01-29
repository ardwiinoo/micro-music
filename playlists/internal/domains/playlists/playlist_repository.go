package playlists

import (
	"github.com/google/uuid"

	"github.com/ardwiinoo/micro-music/playlists/internal/domains/playlists/entities"
)

type PlaylistRepository interface {
	AddPlaylist(playlist *entities.AddPlaylist, userID int) (playlistId uuid.UUID, err error)
}