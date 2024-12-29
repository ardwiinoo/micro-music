package translator

import (
	"github.com/ardwiinoo/micro-music/authentications/internal/commons/exceptions"
)

var DomainErrorMapping = map[string]exceptions.AppError{
	"REGISTER_USER.EMAIL_INVALID": exceptions.InvariantError("email is invalid"),
	"REGISTER_USER.PASSWORD_INVALID_LENGTH": exceptions.InvariantError("password length must be greater than 8 characters"),
	"REGISTER_USER.NOT_CONTAIN_NEEDED_PROPERTY": exceptions.InvariantError("cannot register user because not contain needed property"),
	"LOGIN_USER.EMAIL_INVALID": exceptions.InvariantError("email is invalid"),
	"LOGIN_USER.PASSWORD_INVALID_LENGTH": exceptions.InvariantError("password length must be greater than 8 characters"),
	"LOGIN_USER.NOT_CONTAIN_NEEDED_PROPERTY": exceptions.InvariantError("cannot login user because not contain needed property"),
	"DETAIL_USER.NOT_CONTAIN_NEEDED_PROPERTY": exceptions.InvariantError("cannot create detail user because not contain needed property"),
}
