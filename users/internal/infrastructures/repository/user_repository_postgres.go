package repository

import (
	"github.com/ardwiinoo/micro-music/users/internal/domains/users"
	"github.com/jmoiron/sqlx"
)

type userRepositoryPostgres struct {
	db *sqlx.DB
}

func NewUserRepositoryPostgres(db *sqlx.DB) users.UserRepository {
	return &userRepositoryPostgres{
		db: db,
	}
}
