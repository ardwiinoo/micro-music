package entities

import "github.com/google/uuid"

type SongDetail struct {
	PlaylistID   uuid.UUID `json:"-" db:"playlist_id"`
	PlaylistName string    `json:"-" db:"playlist_name"`
	SongID       uuid.UUID `json:"song_id" db:"song_id"`
	Title        string    `json:"title" db:"title"`
	Artist       string    `json:"artist" db:"artist"`
	Duration     int       `json:"year" db:"year"`
}

type ExportPlaylist struct {
	PlaylistID   uuid.UUID    `json:"playlist_id"`
	PlaylistName string       `json:"playlist_name"`
	Songs        []SongDetail `json:"songs"`
}
