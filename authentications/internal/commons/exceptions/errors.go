package exceptions

import "errors"

var (
	ErrEmailInvalid = errors.New("email is invalid")
	ErrEmailRequired = errors.New("email is required")
	ErrPasswordRequired = errors.New("password is required")
	ErrEmailAlreadyUsed = errors.New("email is already used")
	ErrPasswordInvalidLength = errors.New("password length must be at least 8 characters")
	ErrFullNameRequired = errors.New("full name is required")
	ErrInternalServerError = errors.New("internal server error")
)

var (
	ErrInvalidPaylod = errors.New("invalid payload")
)