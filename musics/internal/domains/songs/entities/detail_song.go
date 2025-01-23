package entities

import (
	"errors"

	"github.com/google/uuid"
)

type DetailSong struct {
	ID     uuid.UUID    `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Year   int    `json:"year" db:"year"`
	Artist string `json:"artist" db:"artist"`
}

func (d *DetailSong) Validate() (err error) {
	if d.ID == uuid.Nil || d.Title == "" || d.Year == 0 || d.Artist == "" {
		return errors.New("DETAIL_SONG.NOT_CONTAIN_NEEDED_PROPERTY")
	}

	return
}