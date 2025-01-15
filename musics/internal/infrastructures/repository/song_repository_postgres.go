package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/ardwiinoo/micro-music/musics/internal/domains/songs"
	"github.com/ardwiinoo/micro-music/musics/internal/domains/songs/entities"
)

type songRepositoryPostgres struct {
	db *sqlx.DB
}


func NewSongRepositoryPostgres(db *sqlx.DB) songs.SongRepository {
	return &songRepositoryPostgres{
		db: db,
	}
}

// AddSong implements songs.SongRepository.
func (s *songRepositoryPostgres) AddSong(ctx context.Context, payload entities.AddSong) (id string, err error) {
	query := `
		INSERT INTO songs 
			(id, title, year, artist, created_at, updated_at)
		VALUES
			(:id, :title, :year, :artist, :created_at, :updated_at)
		RETURNING id
	`

	params := map[string]interface{}{
		"id": uuid.Must(uuid.NewRandom()).String(),
		"title": payload.Title,
		"year": payload.Year,
		"artist": payload.Artist,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}

	stmt, err := s.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()
	
	err = stmt.QueryRowContext(ctx, params).Scan(&id)
	if err != nil {
		return
	}

	return
}

// GetListSong implements songs.SongRepository.
func (s *songRepositoryPostgres) GetListSong(ctx context.Context) (listSong []entities.DetailSong, err error) {
	query := `
		SELECT id, title, year, artist
		FROM songs
	`

	err = s.db.SelectContext(ctx, &listSong, query)
	if err != nil {
		return
	}

	return
}
