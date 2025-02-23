package translator

import "github.com/ardwiinoo/micro-music/users/internal/commons/exceptions"

var DomainErrorMapping = map[string]exceptions.AppError{
	"ADD_USER.NOT_CONTAIN_NEEDED_PROPERTY": exceptions.InvariantError("cannot add user because not contain needed property"),
	"ADD_USER.EMAIL_INVALID":               exceptions.InvariantError("email invalid"),
	"ADD_USER.PASSWORD_INVALID_LENGTH":     exceptions.InvariantError("password length must be greater than 8 characters"),
}
