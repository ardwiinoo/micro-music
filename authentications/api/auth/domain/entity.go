package domain

import (
	"strings"
	"time"

	request "github.com/ardwiinoo/micro-music/authentications/api/auth/app"
	"github.com/ardwiinoo/micro-music/authentications/internal/exception"
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

func NewFromRegisterRequest(req request.RegisterRequestPayload) UserEntity {
	return UserEntity{
		Email:    req.Email,
		FullName: req.FullName,
		Password: req.Password,
		PublicId: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewFromLoginRequest(req request.LoginRequestPayload) UserEntity {
	return UserEntity{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (u UserEntity) validate() (err error) {
	if err = u.validateEmail(); err != nil {
		return
	}

	return
}

func (u UserEntity) validateEmail() (err error) {
	if u.Email == "" {
		return exception.ErrEmailRequired
	}

	emails := strings.Split(u.Email, "@")
	
	if (len(emails) < 2) {
		return exception.ErrEmailInvalid
	}

	return
}