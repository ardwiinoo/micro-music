package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users"
	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users/entities"
)

type userRepositoryPostgres struct {
	db *sqlx.DB
}

func NewUserRepositoryPostgres(db *sqlx.DB) users.UserRepository {
	return &userRepositoryPostgres{
		db: db,
	}
}

// AddUser implements users.UserRepository.
func (u *userRepositoryPostgres) AddUser(ctx context.Context, payload entities.RegisterUser) {
	panic("unimplemented")
}

// VerifyAvailableEmail implements users.UserRepository.
func (u *userRepositoryPostgres) VerifyAvailableEmail(ctx context.Context, email string) {
	panic("unimplemented")
}