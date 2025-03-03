package users

import (
	"context"

	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users/entities"
	"github.com/google/uuid"
)

type UserRepository interface {
	VerifyAvailableEmail(ctx context.Context, email string) (err error)
	AddUser(ctx context.Context, payload entities.RegisterUser) (publicId uuid.UUID, err error)
	GetUserByEmail(ctx context.Context, email string) (user entities.DetailUser, err error)
	ActivateUser(ctx context.Context, publicId uuid.UUID) (err error)
	VerifyUser(ctx context.Context, publicId uuid.UUID) (err error)
}
