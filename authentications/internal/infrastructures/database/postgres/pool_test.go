package postgres_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ardwiinoo/micro-music/authentications/config"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures/database/postgres"
)

func init() {
	if err := config.LoadConfig("../../../../.env"); err != nil {
		panic(err)
	}
}

func TestConnectionPostgres(t *testing.T) {
	originalConfig := config.Cfg

	t.Run("Success Connect Postgres", func(t *testing.T) {
		db, err := postgres.ConnectPostgres()

		require.Nil(t, err)
		require.NotNil(t, db)

		if db != nil {
			db.Close()
		}
	})

	t.Run("Failed Connect Postgres", func(t *testing.T) {
		config.Cfg.DB.Host = "invalid_host"
		config.Cfg.DB.Port = 9999
		config.Cfg.DB.User = "invalid_user"
		config.Cfg.DB.Password = "invalid_password"
		config.Cfg.DB.DBName = "invalid_db"

		db, err := postgres.ConnectPostgres()

		require.NotNil(t, err)
		require.Nil(t, db)

		config.Cfg = originalConfig
	})
}