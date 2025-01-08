package infrastructures

import (
	"github.com/jmoiron/sqlx"

	"github.com/ardwiinoo/micro-music/authentications/config"
	usecase "github.com/ardwiinoo/micro-music/authentications/internal/applications/use_case"
	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures/broker/rabbitmq"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures/database/postgres"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures/repository"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures/security"
)

type Container struct {
	DB *sqlx.DB
	UserRepository   users.UserRepository
	RegisterUserUseCase usecase.RegisterUserUseCase
	LoginUserUseCase usecase.LoginUserUseCase
	RabbitMQ *rabbitmq.RabbitMQ
}

func NewContainer() (container *Container, err error) {
	
	db, err := postgres.ConnectPostgres()
	if err != nil {
		return nil, err
	}

	rabbit, err := rabbitmq.NewRabbitMQ(config.Cfg.Rabbit.ConnString)
	if err != nil {
		db.Close()
		return nil, err
	}

	// Security
	passwordHash := security.NewPasswordHash()
	pasetoManager := security.NewPasetoTokenManager(config.Cfg.App.AppSecret.AppPrivateKey, config.Cfg.App.AppSecret.AppPublicKey)

	// Repository
	userRepo := repository.NewUserRepositoryPostgres(db)

	// UseCase
	registerUserUseCase := usecase.NewRegisterUserUseCase(userRepo, passwordHash, rabbit)
	loginUserUseCase := usecase.NewloginUserUseCase(userRepo, passwordHash, pasetoManager)

	return &Container{
		DB: db,
		RabbitMQ: rabbit,
		UserRepository: userRepo,
		RegisterUserUseCase: registerUserUseCase,
		LoginUserUseCase: loginUserUseCase,
	}, nil
}

func (c *Container) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
	
	if c.RabbitMQ != nil {
		c.RabbitMQ.Close()
	}
}