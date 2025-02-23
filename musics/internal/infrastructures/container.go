package infrastructures

import (
	"github.com/ardwiinoo/micro-music/musics/internal/applications/service"
	"github.com/jmoiron/sqlx"

	"github.com/ardwiinoo/micro-music/musics/config"
	appCache "github.com/ardwiinoo/micro-music/musics/internal/applications/cache"
	appSecurity "github.com/ardwiinoo/micro-music/musics/internal/applications/security"
	usecase "github.com/ardwiinoo/micro-music/musics/internal/applications/use_case"
	"github.com/ardwiinoo/micro-music/musics/internal/domains/songs"
	infraCache "github.com/ardwiinoo/micro-music/musics/internal/infrastructures/cache"
	"github.com/ardwiinoo/micro-music/musics/internal/infrastructures/database/postgres"
	"github.com/ardwiinoo/micro-music/musics/internal/infrastructures/repository"
	infraSecurity "github.com/ardwiinoo/micro-music/musics/internal/infrastructures/security"
	infraService "github.com/ardwiinoo/micro-music/musics/internal/infrastructures/service"
)

type Container struct {
	DB                 *sqlx.DB
	Redis              appCache.CacheManager
	FirebaseStorage    service.FirebaseService
	TokenManager       appSecurity.TokenManager
	SongRepository     songs.SongRepository
	AddSongUseCase     usecase.AddSongUseCase
	GetListSongUseCase usecase.GetListSongUseCase
	StreamSongUseCase  usecase.StreamSongUseCase
}

func NewContainer() (container *Container, err error) {

	// Database
	db, err := postgres.ConnectPostgres()
	if err != nil {
		return nil, err
	}

	// Cache
	redis, err := infraCache.NewRedisCache(config.Cfg.Cache.Host, config.Cfg.Cache.Password, config.Cfg.Cache.DB)
	if err != nil {
		return nil, err
	}

	// Firebase Storage
	firebaseStorage, err := infraService.NewFirebaseStorage(config.Cfg.StorageConfig.CredentialsFile, config.Cfg.StorageConfig.BucketName)
	if err != nil {
		return nil, err
	}

	// Security
	pasetoManager := infraSecurity.NewPasetoTokenManager(config.Cfg.App.AppSecret.AppPublicKey)

	// Repository
	songRepository := repository.NewSongRepositoryPostgres(db)
	userRepository := repository.NewUserRepository(db)

	// Use case
	addSongUseCase := usecase.NewAddSongUseCase(songRepository, userRepository, firebaseStorage, redis)
	getListSongUseCase := usecase.NewGetListSongUseCase(songRepository, redis)
	streamSongUseCase := usecase.NewStreamSongUseCase(songRepository)

	return &Container{
		DB:                 db,
		Redis:              redis,
		FirebaseStorage:    firebaseStorage,
		TokenManager:       pasetoManager,
		SongRepository:     songRepository,
		AddSongUseCase:     addSongUseCase,
		GetListSongUseCase: getListSongUseCase,
		StreamSongUseCase:  streamSongUseCase,
	}, nil
}

func (c *Container) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}
