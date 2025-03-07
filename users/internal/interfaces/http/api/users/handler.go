package users

import (
	"fmt"
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

// AddUserHandler godoc
// @Summary      Add a new user
// @Description  Create a new user in the system
// @Tags         Users
// @Param        Authorization header   string true  "Authorization Bearer Token"
// @Accept       json
// @Produce      json
// @Param        request body entities.AddUser true "Add User Payload"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /users [post]
// @Security     ApiKeyAuth
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

// GetListUserHandler godoc
// @Summary      Get list of users
// @Description  Retrieve a list of users from the system
// @Tags         Users
// @Param        Authorization header   string true  "Authorization Bearer Token"
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /users [get]
// @Security     ApiKeyAuth
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

// DeleteUserHandler godoc
// @Summary      Delete a user
// @Description  Remove a user from the system by ID
// @Tags         Users
// @Param        id   path      int     true  "User ID"
// @Param        Authorization header   string true  "Authorization Bearer Token"
// @Success      204  {object}  nil
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /users/{id} [delete]
// @Security     ApiKeyAuth
func (u *userHandler) DeleteUserHandler(ctx *fiber.Ctx) error {
	userIdStr := ctx.Params("id")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return exceptions.InvariantError("invalid user id")
	}

	fmt.Println(userId)

	err = u.container.DeleteUserUseCase.Execute(ctx.UserContext(), userId)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
