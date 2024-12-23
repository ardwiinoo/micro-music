package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserEntity struct {
	Id        int       `db:"id"`
	Email     string    `db:"email"`
	FullName  string    `db:"full_name"`
	Password  string    `db:"password"`
	PublicId  uuid.UUID `db:"public_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewUserEntity(email, fullName, password string) UserEntity {
	return UserEntity{
		Email:    email,
		FullName: fullName,
		Password: password,
		PublicId: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u UserEntity) IsExists() bool {
	return u.Id != 0
}