package infrastructures

import (
	"github.com/ardwiinoo/micro-music/users/config"
	appSecurity "github.com/ardwiinoo/micro-music/users/internal/applications/security"
	usecase "github.com/ardwiinoo/micro-music/users/internal/applications/use_case"
	"github.com/ardwiinoo/micro-music/users/internal/domains/users"
	"github.com/ardwiinoo/micro-music/users/internal/infrastructures/database/postgres"
	"github.com/ardwiinoo/micro-music/users/internal/infrastructures/repository"
	infraSecurity "github.com/ardwiinoo/micro-music/users/internal/infrastructures/security"
	"github.com/jmoiron/sqlx"
)

type Container struct {
	DB                 *sqlx.DB
	UserRepository     users.UserRepository
	AddUserUseCase     usecase.AddUserUseCase
	GetListUserUseCase usecase.GetListUserUseCase
	TokenManager       appSecurity.TokenManager
}

func NewContainer() (container *Container, err error) {

	db, err := postgres.ConnectPostgres()
	if err != nil {
		return nil, err
	}

	// Security
	passwordHash := infraSecurity.NewPasswordHash()
	pasetoManager := infraSecurity.NewPasetoTokenManager(config.Cfg.App.AppSecret.AppPublicKey)

	// Repository
	userRepo := repository.NewUserRepositoryPostgres(db)

	// UseCase
	addUserUseCase := usecase.NewAddUserUseCase(userRepo, passwordHash)
	getListUserUseCase := usecase.NewGetListUserUseCase(userRepo)

	return &Container{
		DB:                 db,
		UserRepository:     userRepo,
		AddUserUseCase:     addUserUseCase,
		GetListUserUseCase: getListUserUseCase,
		TokenManager:       pasetoManager,
	}, nil
}

func (c *Container) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}
