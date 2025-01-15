package songs

import (
	"context"

	"github.com/ardwiinoo/micro-music/musics/internal/domains/songs/entities"
)

type SongRepository interface {
	GetListSong(ctx context.Context) (listSong []entities.DetailSong, err error)
	AddSong(ctx context.Context, payload entities.AddSong) (id string, err error)
}