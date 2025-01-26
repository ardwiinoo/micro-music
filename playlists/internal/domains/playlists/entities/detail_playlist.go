package entities

import (
	"errors"

	"github.com/google/uuid"
)

type DetailPlaylist struct {
	ID uuid.UUID `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	OwnerID uuid.UUID `json:"owner_id" db:"owner_id"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

func (a *DetailPlaylist) Validate() (err error) {
	if a.ID == uuid.Nil || a.Name == "" || a.OwnerID == uuid.Nil || a.CreatedAt == "" || a.UpdatedAt == "" {
		return errors.New("DETAIL_PLAYLIST.NOT_CONTAIN_NEEDED_PROPERTY")
	}

	return
}