package users

import (
	"context"

	"github.com/ardwiinoo/micro-music/musics/internal/domains/users/entities"
)

type UserRepository interface {
	GetDetailUserByPublicId(ctx context.Context, publicId string) (userDetail entities.DetailUser, err error)
}