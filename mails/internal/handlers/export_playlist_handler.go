package handlers

import (
	"encoding/json"

	"github.com/ardwiinoo/micro-music/mails/internal/email"
)

type ExportPlaylistPayload struct {
	Email      string `json:"email"`
	PlaylistID string `json:"playlist_id"`
	Playlist   string `json:"playlist"`
}

func ExportPlaylistHandler(mailer *email.Mailer) func([]byte) error {
	return func(payload []byte) error {
		var data ExportPlaylistPayload

		if err := json.Unmarshal(payload, &data); err != nil {
			return err
		}

		var playlist struct {
			PlaylistID   string `json:"playlist_id"`
			PlaylistName string `json:"playlist_name"`
			Songs        []struct {
				Title  string `json:"title"`
				Artist string `json:"artist"`
				Year   int    `json:"year"`
			} `json:"songs"`
		}

		if err := json.Unmarshal([]byte(data.Playlist), &playlist); err != nil {
			return err
		}

		return mailer.SendMail(
			data.Email,
			"Your Exported Playlist - MicroMusic",
			"export_playlist.html",
			map[string]interface{}{
				"PlaylistName": playlist.PlaylistName,
				"Songs":        playlist.Songs,
			},
		)
	}
}
