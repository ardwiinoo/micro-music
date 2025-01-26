package infrastructures

import (
	"github.com/jmoiron/sqlx"

	"github.com/ardwiinoo/micro-music/playlists/config"
	appSecurity "github.com/ardwiinoo/micro-music/playlists/internal/applications/security"
	"github.com/ardwiinoo/micro-music/playlists/internal/infrastructures/database/postgres"
	infraSecurity "github.com/ardwiinoo/micro-music/playlists/internal/infrastructures/security"
)

type Container struct {
	DB *sqlx.DB
	TokenManager appSecurity.TokenManager
}

func NewContainer() (container *Container, err error) {

	// Database
	db, err := postgres.ConnectPostgres()
	if err != nil {
		return nil, err
	}
	
	// Security
	pasetoManager := infraSecurity.NewPasetoTokenManager(config.Cfg.App.AppSecret.AppPublicKey)

	return &Container{
		DB: db,
		TokenManager: pasetoManager,
	}, nil
}

func (c *Container) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}