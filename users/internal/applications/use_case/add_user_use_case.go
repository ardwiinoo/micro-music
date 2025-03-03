package usecase

import (
	"context"
	"errors"
	"github.com/ardwiinoo/micro-music/users/internal/applications/security"
	"github.com/ardwiinoo/micro-music/users/internal/commons/constants"
	"github.com/ardwiinoo/micro-music/users/internal/domains/users"
	"github.com/ardwiinoo/micro-music/users/internal/domains/users/entities"
	"github.com/google/uuid"
)

type AddUserUseCase interface {
	Execute(ctx context.Context, payload entities.AddUser) (publicId uuid.UUID, err error)
}

type addUserUseCase struct {
	userRepository users.UserRepository
	passwordHash   security.PasswordHash
}

func (a addUserUseCase) Execute(ctx context.Context, payload entities.AddUser) (publicId uuid.UUID, err error) {
	publicID, ok := ctx.Value(constants.PublicIDContextKey).(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("ADD_USER_USE_CASE.PUBLIC_ID_NOT_FOUND")
	}

	user, err := a.userRepository.GetDetailUserByPublicId(ctx, publicID.String())
	if err != nil {
		return uuid.Nil, errors.New("ADD_USER_USE_CASE.USER_NOT_FOUND")
	}

	if user.Role != constants.AdminRoleID {
		return uuid.Nil, errors.New("ADD_USER_USE_CASE.USER_NOT_AUTHORIZED")
	}

	err = payload.Validate()
	if err != nil {
		return uuid.Nil, err
	}

	err = a.userRepository.VerifyAvailableEmail(ctx, payload.Email)
	if err != nil {
		return uuid.Nil, err
	}

	hashedPassword, err := a.passwordHash.Hash(payload.Password)
	if err != nil {
		return uuid.Nil, err
	}

	payload.Password = hashedPassword
	publicId, err = a.userRepository.AddUser(ctx, payload)
	if err != nil {
		return uuid.Nil, err
	}

	// TODO: Send notif to user email

	return publicId, nil
}

func NewAddUserUseCase(userRepository users.UserRepository, passwordHash security.PasswordHash) AddUserUseCase {
	return &addUserUseCase{
		userRepository: userRepository,
		passwordHash:   passwordHash,
	}
}
