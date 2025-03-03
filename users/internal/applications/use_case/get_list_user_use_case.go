package usecase

import (
	"context"
	"errors"
	"github.com/ardwiinoo/micro-music/users/internal/commons/constants"
	"github.com/ardwiinoo/micro-music/users/internal/domains/users"
	"github.com/ardwiinoo/micro-music/users/internal/domains/users/entities"
	"github.com/google/uuid"
)

type GetListUserUseCase interface {
	Execute(ctx context.Context) ([]entities.DetailUser, error)
}

type getListUserUseCase struct {
	userRepository users.UserRepository
}

func (g getListUserUseCase) Execute(ctx context.Context) ([]entities.DetailUser, error) {
	publicID, ok := ctx.Value(constants.PublicIDContextKey).(uuid.UUID)
	if !ok {
		return nil, errors.New("GET_LIST_USER_USE_CASE.PUBLIC_ID_NOT_FOUND")
	}

	user, err := g.userRepository.GetDetailUserByPublicId(ctx, publicID.String())
	if err != nil {
		return nil, errors.New("GET_LIST_USER_USE_CASE.USER_NOT_FOUND")
	}

	if user.Role != constants.AdminRoleID {
		return nil, errors.New("GET_LIST_USER_USE_CASE.USER_NOT_AUTHORIZED")
	}

	return g.userRepository.GetListUser(ctx)
}

func NewGetListUserUseCase(userRepository users.UserRepository) GetListUserUseCase {
	return &getListUserUseCase{
		userRepository: userRepository,
	}
}
