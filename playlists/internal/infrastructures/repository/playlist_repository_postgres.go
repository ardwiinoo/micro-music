package repository

import (
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
func (p *playlistRepository) AddPlaylist(playlist *entities.AddPlaylist, userID int) (playlistId uuid.UUID, err error) {
	query :=`
		INSERT INTO 
			playlists (id, name, created_at, updated_at, owner_id)
		VALUES
			(:id, :name, :created_at, :updated_at, :owner_id)
		RETURNING id
	`

	params := map[string]interface{}{
		"id": uuid.Must(uuid.NewRandom()),
		"name": playlist.Name,
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"owner_id": userID,
	}

	err = p.db.Get(&playlistId, query, params)
	if err != nil {
		return uuid.Nil, err
	}

	return playlistId, nil
}
