package infrastructures

import (
	"github.com/jmoiron/sqlx"

	"github.com/ardwiinoo/micro-music/authentications/config"
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
	LoginUserUseCase usecase.LoginUserUseCase
}

func NewContainer() (container *Container, err error) {
	
	db, err := postgres.ConnectPostgres()

	if err != nil {
		return nil, err
	}

	passwordHash := security.NewPasswordHash()
	pasetoManager := security.NewPasetoTokenManager(config.Cfg.App.AppSecret.AppPrivateKey, config.Cfg.App.AppSecret.AppPublicKey)

	userRepo := repository.NewUserRepositoryPostgres(db)

	registerUserUseCase := usecase.NewRegisterUserUseCase(userRepo, passwordHash)
	loginUserUseCase := usecase.NewloginUserUseCase(userRepo, passwordHash, pasetoManager)

	return &Container{
		DB: db,
		UserRepository: userRepo,
		RegisterUserUseCase: registerUserUseCase,
		LoginUserUseCase: loginUserUseCase,
	}, nil
}

func (c *Container) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}