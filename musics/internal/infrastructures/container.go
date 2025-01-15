package infrastructures

import (
	"github.com/jmoiron/sqlx"

	"github.com/ardwiinoo/micro-music/musics/config"
	appSecurity "github.com/ardwiinoo/micro-music/musics/internal/applications/security"
	usecase "github.com/ardwiinoo/micro-music/musics/internal/applications/use_case"
	"github.com/ardwiinoo/micro-music/musics/internal/domains/songs"
	"github.com/ardwiinoo/micro-music/musics/internal/infrastructures/database/postgres"
	"github.com/ardwiinoo/micro-music/musics/internal/infrastructures/repository"
	infraSecurity "github.com/ardwiinoo/micro-music/musics/internal/infrastructures/security"
)

type Container struct {
	DB *sqlx.DB
	TokenManager appSecurity.TokenManager
	SongRepository songs.SongRepository
	AddSongUseCase usecase.AddSongUseCase
	GetListSongUseCase usecase.GetListSongUseCase

}

func NewContainer() (container *Container, err error) {
	
	// Database
	db, err := postgres.ConnectPostgres()
	if err != nil {
		return nil, err
	}

	// Security
	pasetoManager := infraSecurity.NewPasetoTokenManager(config.Cfg.App.AppSecret.AppPublicKey)

	// Repository
	songRepository := repository.NewSongRepositoryPostgres(db)

	// Use case
	addSongUseCase := usecase.NewAddSongUseCase(songRepository)
	getListSongUseCase := usecase.NewGetListSongUseCase(songRepository)

	return &Container{
		DB: db,
		TokenManager: pasetoManager,
		SongRepository: songRepository,
		AddSongUseCase: addSongUseCase,
		GetListSongUseCase: getListSongUseCase,
	}, nil
}

func (c *Container) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}