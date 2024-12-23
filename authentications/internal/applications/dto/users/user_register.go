package dto

import (
	"strings"

	"github.com/ardwiinoo/micro-music/authentications/internal/commons/exceptions"
	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users/entities"
)

type UserRegister struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

func NewUserRegister(email, fullName, password string) (userDto UserRegister, err error) {

	dto := UserRegister{
		Email:    email,
		FullName: fullName,
		Password: password,
	}

	if err := dto.Validate(); err != nil {
		return UserRegister{}, err
	}

	return
}

func (u UserRegister) Validate() (err error) {
	if err = u.ValidateEmail(); err != nil {
		return
	}

	if err = u.ValidatePassword(); err != nil {
		return
	}

	if err = u.ValidateFullName(); err != nil {
		return
	}

	return
}

func (u UserRegister) ValidateEmail() (err error) {
	if u.Email == "" {
		return exceptions.ErrEmailRequired
	}

	parts := strings.Split(u.Email, "@")
	if len(parts) != 2 {
		return exceptions.ErrEmailInvalid
	}

	return nil
}

func (u UserRegister) ValidatePassword() (err error) {
	if len(u.Password) < 8 {
		return exceptions.ErrPasswordInvalidLength
	}

	return nil
}

func (u UserRegister) ValidateFullName() (err error) {
	if u.FullName == "" {
		return exceptions.ErrFullNameRequired
	}

	return nil
}

func (u UserRegister) ToEntity() entities.UserEntity {
	return entities.NewUserEntity(u.Email, u.FullName, u.Password)
}