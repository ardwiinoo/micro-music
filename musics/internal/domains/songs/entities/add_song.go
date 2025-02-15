package entities

import "errors"

type AddSong struct {
	Title  string `json:"title"`
	Year   int    `json:"year"`
	Artist string `json:"artist"`
	Url    string `json:"url"`
}

func (a *AddSong) Validate() (err error) {
	if a.Title == "" || a.Year == 0 || a.Artist == "" {
		return errors.New("ADD_SONG.NOT_CONTAIN_NEEDED_PROPERTY")
	}

	return
}
