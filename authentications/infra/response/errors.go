package response

import (
	"net/http"

	"github.com/ardwiinoo/micro-music/authentications/internal/exception"
)

type Error struct {
	Message  string `json:"message"`
	Code     string `json:"code"`
	HttpCode int    `json:"http_code"`
}

func NewError(message string, code string, httpCode int) Error {
	return Error{
		Message:  message,
		Code:     code,
		HttpCode: httpCode,
	}
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrorEmailRequired    = NewError(exception.ErrEmailRequired.Error(), "40001", http.StatusBadRequest)
	ErrorPasswordRequired = NewError(exception.ErrPasswordRequired.Error(), "40002", http.StatusBadRequest)
	ErrorEmailAlreadyUsed = NewError(exception.ErrEmailAlreadyUsed.Error(), "40901", http.StatusConflict)
	ErrorEmailInvalid     = NewError(exception.ErrEmailInvalid.Error(), "40003", http.StatusBadRequest)
)

var ErrorMapping = map[string]Error{
	exception.ErrEmailRequired.Error(): ErrorEmailRequired,
	exception.ErrPasswordRequired.Error(): ErrorPasswordRequired,
	exception.ErrEmailAlreadyUsed.Error(): ErrorEmailAlreadyUsed,
	exception.ErrEmailInvalid.Error(): ErrorEmailInvalid,
}