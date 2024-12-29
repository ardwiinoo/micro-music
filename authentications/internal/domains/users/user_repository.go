package users

import (
	"context"

	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users/entities"
)

type UserRepository interface {
	VerifyAvailableEmail(ctx context.Context, email string) (err error)
	AddUser(ctx context.Context, payload entities.RegisterUser) (id int, err error)
	GetUserByEmail(ctx context.Context, email string) (user entities.DetailUser, err error)
}