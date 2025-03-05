package users

import (
	"github.com/ardwiinoo/micro-music/users/internal/commons/exceptions"
	"github.com/ardwiinoo/micro-music/users/internal/domains/users/entities"
	"github.com/ardwiinoo/micro-music/users/internal/infrastructures"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type userHandler struct {
	container infrastructures.Container
}

func NewUserHandler(container infrastructures.Container) *userHandler {
	return &userHandler{
		container: container,
	}
}

func (u *userHandler) AddUserHandler(ctx *fiber.Ctx) error {
	var payload = entities.AddUser{}

	if err := ctx.BodyParser(&payload); err != nil {
		return exceptions.InvariantError("invalid payload")
	}

	publicId, err := u.container.AddUserUseCase.Execute(ctx.UserContext(), payload)

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"publicId": publicId,
		},
	})
}

func (u *userHandler) GetListUserHandler(ctx *fiber.Ctx) error {
	listUser, err := u.container.GetListUserUseCase.Execute(ctx.UserContext())
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"data":   listUser,
	})
}

func (u *userHandler) DeleteUserHandler(ctx *fiber.Ctx) error {
	userIdStr := ctx.Params("id")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return exceptions.InvariantError("invalid user id")
	}

	err = u.container.DeleteUserUseCase.Execute(ctx.UserContext(), userId)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
