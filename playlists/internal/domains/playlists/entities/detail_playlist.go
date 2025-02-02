package entities

import (
	"errors"

	"github.com/google/uuid"
)

type DetailPlaylist struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Owner     string    `json:"owner" db:"owner"`
	CreatedAt string    `json:"created_at" db:"created_at"`
	UpdatedAt string    `json:"updated_at" db:"updated_at"`
}

func (a *DetailPlaylist) Validate() (err error) {
	if a.ID == uuid.Nil || a.Name == "" || a.Owner == "" || a.CreatedAt == "" || a.UpdatedAt == "" {
		return errors.New("DETAIL_PLAYLIST.NOT_CONTAIN_NEEDED_PROPERTY")
	}

	return
}
