package infrastructures

import (
	"github.com/ardwiinoo/micro-music/users/config"
	"github.com/ardwiinoo/micro-music/users/internal/domains/users"
	"github.com/ardwiinoo/micro-music/users/internal/infrastructures/database/postgres"
	"github.com/ardwiinoo/micro-music/users/internal/infrastructures/repository"
	"github.com/ardwiinoo/micro-music/users/internal/infrastructures/security"
	"github.com/jmoiron/sqlx"
)

type Container struct {
	DB             *sqlx.DB
	UserRepository users.UserRepository
}

func NewContainer() (container *Container, err error) {

	db, err := postgres.ConnectPostgres()
	if err != nil {
		return nil, err
	}

	// Security
	passwordHash := security.NewPasswordHash()
	pasetoManager := security.NewPasetoTokenManager(config.Cfg.App.AppSecret.AppPublicKey)

	// Repository
	userRepo := repository.NewUserRepositoryPostgres(db)

	return &Container{
		DB:             db,
		UserRepository: userRepo,
	}, nil
}

func (c *Container) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}
