package infrastructures

import (
	"github.com/jmoiron/sqlx"

	"github.com/ardwiinoo/micro-music/playlists/config"
	appSecurity "github.com/ardwiinoo/micro-music/playlists/internal/applications/security"
	usecase "github.com/ardwiinoo/micro-music/playlists/internal/applications/use_case"
	"github.com/ardwiinoo/micro-music/playlists/internal/infrastructures/broker/rabbitmq"
	"github.com/ardwiinoo/micro-music/playlists/internal/infrastructures/database/postgres"
	"github.com/ardwiinoo/micro-music/playlists/internal/infrastructures/repository"
	infraSecurity "github.com/ardwiinoo/micro-music/playlists/internal/infrastructures/security"
)

type Container struct {
	DB                     *sqlx.DB
	TokenManager           appSecurity.TokenManager
	AddPlaylistUseCase     usecase.AddPlaylistUseCase
	DeletePlaylistUseCase  usecase.DeletePlaylistUseCase
	AddPlaylistSongUseCase usecase.AddPlaylistSongUseCase
	GetListPlaylistUseCase usecase.GetListPlaylistUseCase
	ExportPlaylistUseCase  usecase.ExportPlaylistUseCase
	RabbitMQ               *rabbitmq.RabbitMQ
}

func NewContainer() (container *Container, err error) {

	// Database
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
	pasetoManager := infraSecurity.NewPasetoTokenManager(config.Cfg.App.AppSecret.AppPublicKey)

	// Repository
	userRepository := repository.NewUserRepository(db)
	playlistRepository := repository.NewPlaylistRepository(db)

	// UseCase
	addPlaylistUseCase := usecase.NewAddPlaylistUseCase(playlistRepository, userRepository)
	deletePlaylistUseCase := usecase.NewDeletePlaylistUseCase(playlistRepository, userRepository)
	addPlaylistSongUseCase := usecase.NewAddPlaylistSongUseCase(playlistRepository, userRepository)
	getListPlaylistUseCase := usecase.NewGetListPlaylistUseCase(playlistRepository, userRepository)
	exportPlaylistUseCase := usecase.NewExportPlaylistUseCase(playlistRepository, userRepository, rabbit)

	return &Container{
		DB:                     db,
		RabbitMQ:               rabbit,
		TokenManager:           pasetoManager,
		AddPlaylistUseCase:     addPlaylistUseCase,
		DeletePlaylistUseCase:  deletePlaylistUseCase,
		AddPlaylistSongUseCase: addPlaylistSongUseCase,
		GetListPlaylistUseCase: getListPlaylistUseCase,
		ExportPlaylistUseCase:  exportPlaylistUseCase,
	}, nil
}

func (c *Container) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}
