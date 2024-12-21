package exception

import "errors"

var (
	ErrEmailRequired    = errors.New("email is required")
	ErrPasswordRequired = errors.New("password is required")
	ErrEmailAlreadyUsed = errors.New("email already used")
	ErrEmailInvalid	 = errors.New("email is invalid")
)