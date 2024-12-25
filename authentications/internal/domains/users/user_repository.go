package users

import (
	"context"

	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users/entities"
)

type UserRepository interface {
	VerifyAvailableEmail(ctx context.Context, email string)
	AddUser(ctx context.Context, payload entities.RegisterUser)
}