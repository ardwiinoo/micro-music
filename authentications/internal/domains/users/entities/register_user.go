package entities

import (
	"regexp"

	"github.com/ardwiinoo/micro-music/authentications/internal/commons/exceptions"
)

type RegisterUser struct {
	FullName string
	Email    string
	Password string
}

func (r *RegisterUser) Validate() (err error) {
	if r.FullName == "" {
		return exceptions.ErrFullNameRequired
	}

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

func isValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}