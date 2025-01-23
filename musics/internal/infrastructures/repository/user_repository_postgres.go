package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/ardwiinoo/micro-music/musics/internal/domains/users"
	"github.com/ardwiinoo/micro-music/musics/internal/domains/users/entities"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) users.UserRepository {
	return &userRepository{
		db: db,
	}
}

// GetDetailUserByPublicId implements users.UserRepository.
func (u *userRepository) GetDetailUserByPublicId(ctx context.Context, publicId string) (userDetail entities.DetailUser, err error) {
    query := `
        SELECT
            id, email, full_name, public_id, role, created_at, updated_at
        FROM
            users
        WHERE
            public_id = $1
    `

    err = u.db.GetContext(ctx, &userDetail, query, publicId)
    if err != nil {
        return
    }

    return
}

