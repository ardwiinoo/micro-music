package entities

import (
	"github.com/ardwiinoo/micro-music/authentications/internal/commons/exceptions"
)


type LoginUser struct {
	Email    string
	Password string
}

func (r *LoginUser) Validate() (err error) {
	if r.Email == "" {
		return exceptions.ErrEmailRequired
	}

	if !isValidEmail(r.Email) {
		return exceptions.ErrEmailInvalid
	}

	if len(r.Password) < 8 {
		return exceptions.ErrPasswordInvalidLength
	}

	return
}