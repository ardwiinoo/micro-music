package usecase

import (
	"context"
	"errors"
	"github.com/ardwiinoo/micro-music/users/internal/domains/users"
)

type DeleteUserUseCase interface {
	Execute(ctx context.Context, userId int) error
}

type deleteUserUseCase struct {
	userRepository users.UserRepository
}

func (d deleteUserUseCase) Execute(ctx context.Context, userId int) error {
	if userId < 1 {
		return errors.New("DELETE_USER_USE_CASE.INVALID_USER_ID")
	}

	return d.userRepository.DeleteUserById(ctx, userId)
}

func NewDeleteUserUseCase(userRepository users.UserRepository) DeleteUserUseCase {
	return &deleteUserUseCase{
		userRepository: userRepository,
	}
}
