package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/ardwiinoo/micro-music/playlists/internal/domains/playlists"
	"github.com/ardwiinoo/micro-music/playlists/internal/domains/playlists/entities"
)

type playlistRepository struct {
	db *sqlx.DB
}

func NewPlaylistRepository(db *sqlx.DB) playlists.PlaylistRepository {
	return &playlistRepository{
		db: db,
	}
}

// AddPlaylist implements playlists.PlaylistRepository.
func (p *playlistRepository) AddPlaylist(ctx context.Context, playlist *entities.AddPlaylist, userID int) (playlistId uuid.UUID, err error) {
	query := `
		INSERT INTO 
			playlists (id, name, created_at, updated_at, owner_id)
		VALUES
			(:id, :name, :created_at, :updated_at, :owner_id)
		RETURNING id
	`

	params := map[string]interface{}{
		"id":         uuid.Must(uuid.NewRandom()),
		"name":       playlist.Name,
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"owner_id":   userID,
	}

	rows, err := p.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return uuid.Nil, err
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	if !rows.Next() {
		return uuid.Nil, sql.ErrNoRows
	}

	err = rows.Scan(&playlistId)
	if err != nil {
		return uuid.Nil, err
	}

	return playlistId, nil
}

// DeletePlaylist implements playlists.PlaylistRepository.
func (p *playlistRepository) DeletePlaylist(ctx context.Context, playlistID uuid.UUID) (err error) {
	query := `
		DELETE FROM
			playlists
		WHERE
			id = $1
	`

	_, err = p.db.ExecContext(ctx, query, playlistID)
	if err != nil {
		return err
	}

	return nil
}

// ValidatePlaylistOwner implements playlists.PlaylistRepository.
func (p *playlistRepository) ValidatePlaylistOwner(ctx context.Context, playlistID uuid.UUID, userID int) (err error) {
	query := `
		SELECT
			id
		FROM
			playlists
		WHERE
			id = $1 AND owner_id = $2
	`

	var id uuid.UUID
	err = p.db.GetContext(ctx, &id, query, playlistID, userID)
	if err != nil {
		return err
	}

	return nil
}

// AddPlaylistSong implements playlists.PlaylistRepository.
func (p *playlistRepository) AddPlaylistSong(ctx context.Context, playlistID uuid.UUID, songID uuid.UUID) (err error) {
	query := `
		INSERT INTO
			playlist_songs (playlist_id, song_id)
		VALUES
			($1, $2)
	`

	_, err = p.db.ExecContext(ctx, query, playlistID, songID)
	if err != nil {
		return err
	}

	return nil
}
