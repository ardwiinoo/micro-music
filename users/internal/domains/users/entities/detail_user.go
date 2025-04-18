package entities

import (
	"errors"
	"github.com/google/uuid"
)

type DetailUser struct {
	ID        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	FullName  string    `json:"full_name" db:"full_name"`
	PublicId  uuid.UUID `json:"public_id" db:"public_id"`
	Role      int       `json:"role" db:"role"`
	IsActive  bool      `json:"is_active,omitempty" db:"is_active"`
	CreatedAt string    `json:"created_at" db:"created_at"`
	UpdatedAt string    `json:"updated_at" db:"updated_at"`
}

func (d *DetailUser) Validate() (err error) {
	if d.ID == 0 || d.Email == "" || d.FullName == "" || d.PublicId == uuid.Nil || d.Role == 0 || d.CreatedAt == "" || d.UpdatedAt == "" {
		return errors.New("DETAIL_USER.NOT_CONTAIN_NEEDED_PROPERTY")
	}

	return
}
