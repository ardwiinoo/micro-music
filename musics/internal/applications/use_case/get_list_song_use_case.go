package usecase

import (
	"context"

	"github.com/ardwiinoo/micro-music/musics/internal/domains/songs"
	"github.com/ardwiinoo/micro-music/musics/internal/domains/songs/entities"
)

type GetListSongUseCase interface {
	Execute(ctx context.Context) (listSong []entities.DetailSong, err error)
}

type getListSongUseCase struct {
	songRepository songs.SongRepository
}

func NewGetListSongUseCase(songRepository songs.SongRepository) GetListSongUseCase {
	return &getListSongUseCase{
		songRepository: songRepository,
	}
}

func (g *getListSongUseCase) Execute(ctx context.Context) (listSong []entities.DetailSong, err error) {
	listSong, err = g.songRepository.GetListSong(ctx)
	if err != nil {
		return nil, err
	}

	return listSong, nil
}