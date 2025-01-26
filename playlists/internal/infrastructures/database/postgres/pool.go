package postgres

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/ardwiinoo/micro-music/playlists/config"
)

func ConnectPostgres() (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Cfg.DB.Host,
		config.Cfg.DB.Port,
		config.Cfg.DB.User,
		config.Cfg.DB.Password,
		config.Cfg.DB.DBName,
	)

	db, err = sqlx.Open("postgres", dsn)

	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		db = nil
		return
	}

	db.SetConnMaxIdleTime(time.Duration(config.Cfg.DB.ConnectionPool.MaxIdletimeConnection) * time.Second)
	db.SetConnMaxLifetime(time.Duration(config.Cfg.DB.ConnectionPool.MaxLifetimeConnection) * time.Second)
	db.SetMaxIdleConns(int(config.Cfg.DB.ConnectionPool.MaxIdleConnection))
	db.SetMaxOpenConns(int(config.Cfg.DB.ConnectionPool.MaxOpenConnection))

	return
}