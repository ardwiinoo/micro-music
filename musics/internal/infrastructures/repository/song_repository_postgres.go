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
			(id, title, year, artist, url, created_at, updated_at)
		VALUES
			(:id, :title, :year, :artist, :url, :created_at, :updated_at)
		RETURNING id
	`

	params := map[string]interface{}{
		"id":         uuid.Must(uuid.NewRandom()).String(),
		"title":      payload.Title,
		"year":       payload.Year,
		"artist":     payload.Artist,
		"url":        payload.Url, // ✅ Tambahkan `url`
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}

	stmt, err := s.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, params).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

// GetListSong implements songs.SongRepository.
func (s *songRepositoryPostgres) GetListSong(ctx context.Context) (listSong []entities.DetailSong, err error) {
	query := `
		SELECT id, title, year, artist, url
		FROM songs
	`

	err = s.db.SelectContext(ctx, &listSong, query)
	if err != nil {
		return
	}

	return
}

// GetSongById implements songs.SongRepository.
func (s *songRepositoryPostgres) GetSongById(ctx context.Context, songId string) (song entities.DetailSong, err error) {
	query := `
		SELECT id, title, year, artist, url
		FROM songs
		WHERE id = $1
	`

	err = s.db.GetContext(ctx, &song, query, songId)
	if err != nil {
		return
	}

	return
}
