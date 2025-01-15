package usecase

import (
	"context"

	"github.com/ardwiinoo/micro-music/musics/internal/domains/songs"
	"github.com/ardwiinoo/micro-music/musics/internal/domains/songs/entities"
)

type AddSongUseCase interface {
	Execute(ctx context.Context, payload entities.AddSong) (id string, err error)
}

type addSongUseCase struct {
	songRepository songs.SongRepository
}

func NewAddSongUseCase(songRepository songs.SongRepository) AddSongUseCase {
	return &addSongUseCase{
		songRepository: songRepository,
	}
}

func (a *addSongUseCase) Execute(ctx context.Context, payload entities.AddSong) (id string, err error) {
	
	
	id, err = a.songRepository.AddSong(ctx, payload)
	if err != nil {
		return "", err
	}

	return id, nil
}