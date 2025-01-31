package entities

import (
	"errors"
	"github.com/google/uuid"
)

type AddPlaylistSong struct {
	SongID uuid.UUID `json:"song_id"`
}

func (a *AddPlaylistSong) Validate() (err error) {
	if a.SongID == uuid.Nil {
		return errors.New("ADD_PLAYLIST_SONG.NOT_CONTAIN_NEEDED_PROPERTY")
	}

	return
}
