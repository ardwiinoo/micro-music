package users

import (
	"context"

	"github.com/ardwiinoo/micro-music/playlists/internal/domains/users/entities"
)

type UserRepository interface {
	GetDetailUserByPublicId(ctx context.Context, publicId string) (userDetail entities.DetailUser, err error)
}