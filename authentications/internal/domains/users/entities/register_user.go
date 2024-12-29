package entities

import (
	"errors"
	"regexp"
)

type RegisterUser struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *RegisterUser) Validate() (err error) {
	if r.FullName == ""  || r.Email == "" || r.Password == "" {
		return errors.New("REGISTER_USER.NOT_CONTAIN_NEEDED_PROPERTY")
	}

	if !isValidEmail(r.Email) {
		return errors.New("REGISTER_USER.EMAIL_INVALID")
	}

	if len(r.Password) < 8 {
		return errors.New("REGISTER_USER.PASSWORD_INVALID_LENGTH")
	}

	return 
}

func isValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}