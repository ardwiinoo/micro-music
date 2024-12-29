package entities

import (
	"errors"
)

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginUser) Validate() (err error) {
	if r.Email == "" || r.Password == "" {
		return errors.New("LOGIN_USER.NOT_CONTAIN_NEEDED_PROPERTY")
	}

	if !isValidEmail(r.Email) {
		return errors.New("LOGIN_USER.EMAIL_INVALID")
	}

	if len(r.Password) < 8 {
		return errors.New("LOGIN_USER.PASSWORD_INVALID_LENGTH")
	}

	return
}
