package infrastructures

import (
	"github.com/jmoiron/sqlx"

	usecase "github.com/ardwiinoo/micro-music/authentications/internal/applications/use_case"
	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures/database/postgres"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures/repository"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures/security"
)

type Container struct {
	DB *sqlx.DB
	UserRepository   users.UserRepository
	RegisterUserUseCase usecase.RegisterUserUseCase
}

func NewContainer() (container *Container, err error) {
	
	db, err := postgres.ConnectPostgres()

	if err != nil {
		return nil, err
	}

	passwordHash := security.NewPasswordHash()

	userRepo := repository.NewUserRepositoryPostgres(db)

	registerUserUseCase := usecase.NewRegisterUserUseCase(userRepo, passwordHash)

	return &Container{
		DB: db,
		UserRepository: userRepo,
		RegisterUserUseCase: registerUserUseCase,
	}, nil
}

func (c *Container) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}