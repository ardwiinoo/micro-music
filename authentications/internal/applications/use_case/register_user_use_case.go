package usecase

import (
	"context"

	"github.com/ardwiinoo/micro-music/authentications/internal/applications/security"
	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users"
	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users/entities"
)

type RegisterUserUseCase interface {
	Execute(ctx context.Context, request entities.RegisterUser) (userId int, err error)
}

type registerUserUseCase struct {
	userRepository users.UserRepository
	passwordHash security.PasswordHash
}

func NewRegisterUserUseCase(userRepository users.UserRepository, passwordHash security.PasswordHash) RegisterUserUseCase {
	return &registerUserUseCase{
		userRepository: userRepository,
		passwordHash: passwordHash,
	}
}

func (r *registerUserUseCase) Execute(ctx context.Context, payload entities.RegisterUser) (userId int, err error) {

	err = payload.Validate()
	if err != nil {
		return 0, err
	}

	err = r.userRepository.VerifyAvailableEmail(ctx, payload.Email)
	if err != nil {
		return 0, err
	}

	hashedPassword, err := r.passwordHash.Hash(payload.Password)
	if err != nil {
		return 0, err
	}

	payload.Password = hashedPassword
	id, err := r.userRepository.AddUser(ctx, payload)
	if err != nil {
		return 0, err
	}

	return id, nil

}