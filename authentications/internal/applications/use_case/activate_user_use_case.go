package usecase

import (
	"context"

	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users"
	"github.com/google/uuid"
)

type ActivateUserUseCase interface {
	Execute(ctx context.Context, publicId uuid.UUID) (err error)
}

type activateUserUseCase struct {
	userRepository users.UserRepository
}

func NewActivateUserUseCase(userRepository users.UserRepository) ActivateUserUseCase {
	return &activateUserUseCase{
		userRepository: userRepository,
	}
}

func (a *activateUserUseCase) Execute(ctx context.Context, publicId uuid.UUID) (err error) {

	err = a.userRepository.VeryfyUser(ctx, publicId)
	if err != nil {
		return
	}

	err = a.userRepository.ActivateUser(ctx, publicId)
	if err != nil {
		return
	}

	return
}
