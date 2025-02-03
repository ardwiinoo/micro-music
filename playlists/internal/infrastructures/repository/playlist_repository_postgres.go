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

func (p *playlistRepository) GetListPlaylistByUserID(ctx context.Context, userID int) (playlists []entities.DetailPlaylist, err error) {
	query := `
		SELECT
			p.id, p.name, p.created_at, p.updated_at, u.full_name AS owner
		FROM
			playlists p
		JOIN
			users u ON p.owner_id = u.id
		WHERE
			p.owner_id = $1
	`

	err = p.db.SelectContext(ctx, &playlists, query, userID)
	if err != nil {
		return nil, err
	}

	return playlists, nil
}

func (p *playlistRepository) GetPlaylistWithSongs(ctx context.Context, playlistID uuid.UUID, userID int) (playlist *entities.ExportPlaylist, err error) {
	query := `
		SELECT 
			p.id AS playlist_id, 
			p.name AS playlist_name, 
			m.id AS song_id, 
			m.title, 
			m.artist, 
			m.year
		FROM  playlists p
		JOIN playlist_songs ps ON p.id = ps.playlist_id
		JOIN songs m ON ps.song_id = m.id
		WHERE p.id = $1 AND p.owner_id = $2;
	`

	var result []entities.SongDetail
	err = p.db.SelectContext(ctx, &result, query, playlistID, userID)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	}

	playlist = &entities.ExportPlaylist{
		PlaylistID:   result[0].PlaylistID,
		PlaylistName: result[0].PlaylistName,
		Songs:        result,
	}

	return playlist, nil
}
