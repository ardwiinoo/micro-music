package translator

import (
	"net/http"

	"github.com/ardwiinoo/micro-music/authentications/internal/commons/exceptions"
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
	ErrorEmailRequired    = NewError(exceptions.ErrEmailRequired.Error(), "40001", http.StatusBadRequest)
	ErrorPasswordRequired = NewError(exceptions.ErrPasswordRequired.Error(), "40002", http.StatusBadRequest)
	ErrorEmailAlreadyUsed = NewError(exceptions.ErrEmailAlreadyUsed.Error(), "40901", http.StatusConflict)
	ErrorEmailInvalid     = NewError(exceptions.ErrEmailInvalid.Error(), "40003", http.StatusBadRequest)
	ErrorPasswordInvalidLength = NewError(exceptions.ErrPasswordInvalidLength.Error(), "40004", http.StatusBadRequest)

	ErrorInternalServer = NewError("internal server error", "50000", http.StatusInternalServerError)
	ErrorInvalidPayload = NewError("invalid payload", "40000", http.StatusBadRequest)
)

var ErrorMapping = map[string]Error{
	exceptions.ErrEmailRequired.Error(): ErrorEmailRequired,
	exceptions.ErrPasswordRequired.Error(): ErrorPasswordRequired,
	exceptions.ErrEmailAlreadyUsed.Error(): ErrorEmailAlreadyUsed,
	exceptions.ErrEmailInvalid.Error(): ErrorEmailInvalid,
	exceptions.ErrInternalServerError.Error(): ErrorInternalServer,
	exceptions.ErrInvalidPaylod.Error(): ErrorInvalidPayload,
	exceptions.ErrPasswordInvalidLength.Error(): ErrorPasswordInvalidLength,
}