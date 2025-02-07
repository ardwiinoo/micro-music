package entities

import (
	"errors"
	"regexp"
)

type AddUser struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *AddUser) Validate() (err error) {
	if a.FullName == "" || a.Email == "" || a.Password == "" {
		return errors.New("ADD_USER.NOT_CONTAIN_NEEDED_PROPERTY")
	}

	if !isValidEmail(a.Email) {
		return errors.New("ADD_USER.EMAIL_INVALID")
	}

	if len(a.Password) < 8 {
		return errors.New("ADD_USER.PASSWORD_INVALID_LENGTH")
	}

	return
}

func isValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
