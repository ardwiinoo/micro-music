package exceptions

import "net/http"

type AppError struct {
	Message  string
	HttpCode int
}

func (e AppError) Error() string {
	return e.Message
}

func InternalServerError(message string) AppError {
	return AppError{
		Message:  message,
		HttpCode: http.StatusInternalServerError,
	}
}

func InvariantError(message string) AppError {
	return AppError{
		Message:  message,
		HttpCode: http.StatusBadRequest,
	}
}

func NotFoundError(message string) AppError {
	return AppError{
		Message:  message,
		HttpCode: http.StatusNotFound,
	}
}