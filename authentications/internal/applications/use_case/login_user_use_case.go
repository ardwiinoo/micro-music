package usecase

import (
	"context"
	"time"

	"github.com/ardwiinoo/micro-music/authentications/internal/applications/security"
	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users"
	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users/entities"
)

type LoginUserUseCase interface {
	Execute(ctx context.Context, request entities.LoginUser) (token string, err error)
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

func (l *loginUserUseCase) Execute(ctx context.Context, request entities.LoginUser) (token string, err error) {
	payload := map[string]interface{}{
		"user_id": "12345",
		"role":    "admin",
	}

	token, err = l.tokenManager.GenerateToken(payload, time.Minute*5)

	return
}
