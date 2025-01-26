package entities

import "errors"

type AddPlaylist struct {
	Name string `json:"name"`
}

func (a *AddPlaylist) Validate() (err error) {
	if a.Name == "" {
		return errors.New("ADD_PLAYLIST.NOT_CONTAIN_NEEDED_PROPERTY")
	}

	return
}