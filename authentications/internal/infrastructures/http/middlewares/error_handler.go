package middlewares

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ardwiinoo/micro-music/authentications/internal/commons/exceptions"
	"github.com/ardwiinoo/micro-music/authentications/internal/commons/exceptions/translator"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var appErr exceptions.AppError

	if e, ok := err.(exceptions.AppError); ok {
		appErr = e
	} else {
		if translatedErr, exists := translator.DomainErrorMapping[err.Error()]; exists {
			appErr = translatedErr
		} else {
			appErr = exceptions.InternalServerError(err.Error()) 
		}
	}

	status := "error" 
	if appErr.HttpCode >= 400 && appErr.HttpCode < 500 {
		status = "fail"
	}

	return ctx.Status(appErr.HttpCode).JSON(fiber.Map{
		"status": status,
		"http_code": appErr.HttpCode,
		"message": appErr.Message, 
	})
}