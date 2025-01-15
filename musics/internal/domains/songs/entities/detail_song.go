package entities

import "errors"

type DetailSong struct {
	ID     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Year   int    `json:"year" db:"year"`
	Artist string `json:"artist" db:"artist"`
}

func (d *DetailSong) Validate() (err error) {
	if d.ID == 0 || d.Title == "" || d.Year == 0 || d.Artist == "" {
		return errors.New("DETAIL_SONG.NOT_CONTAIN_NEEDED_PROPERTY")
	}

	return
}