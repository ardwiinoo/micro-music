package usecase

import (
	"context"
	"time"

	"github.com/ardwiinoo/micro-music/authentications/internal/applications/security"
	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users"
	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users/entities"
)

type LoginUserUseCase interface {
	Execute(ctx context.Context, payload entities.LoginUser) (token string, err error)
}

type loginUserUseCase struct {
	userRepository users.UserRepository
	passwordHash   security.PasswordHash
	tokenManager   security.TokenManager
}

func NewloginUserUseCase(userRepository users.UserRepository, passwordHash security.PasswordHash, tokenManager security.TokenManager) LoginUserUseCase {
	return &loginUserUseCase{
		userRepository: userRepository,
		passwordHash:   passwordHash,
		tokenManager:   tokenManager,
	}
}

func (l *loginUserUseCase) Execute(ctx context.Context, payload entities.LoginUser) (token string, err error) {
	
	err = payload.Validate()
	if err != nil {
		return "", err
	}

	user, err := l.userRepository.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		return "", err
	}

	err = l.passwordHash.Compare(payload.Password, user.Password)
	if err != nil {
		return "", err
	}

	tokenPayload := map[string]interface{}{
		"public_id": user.PublicId,
	}

	token, err = l.tokenManager.GenerateToken(tokenPayload, time.Hour*5)
	if err != nil {
		return "", err
	}

	return token, nil
}
