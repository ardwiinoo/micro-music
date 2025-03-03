package users

import (
	"context"
	"github.com/ardwiinoo/micro-music/users/internal/domains/users/entities"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetListUser(ctx context.Context) (users []entities.DetailUser, err error)
	AddUser(ctx context.Context, payload entities.AddUser) (publicId uuid.UUID, err error)
	VerifyAvailableEmail(ctx context.Context, email string) (err error)
	GetDetailUserByPublicId(ctx context.Context, publicId string) (userDetail entities.DetailUser, err error)
}
