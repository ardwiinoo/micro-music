package usecase

import (
	"context"
	"github.com/ardwiinoo/micro-music/musics/internal/domains/songs"
	"io"
)

type StreamMusicUseCase interface {
	Execute(ctx context.Context, songId string, rangeHeader string) (io.ReadCloser, int, string, error)
}

type streamMusicUseCase struct {
	songRepository songs.SongRepository
}

func NewStreamMusicUseCase(songRepository songs.SongRepository) StreamMusicUseCase {
	return &streamMusicUseCase{
		songRepository: songRepository,
	}
}

func (s streamMusicUseCase) Execute(ctx context.Context, songId string, rangeHeader string) (io.ReadCloser, int, string, error) {
	//TODO implement me
	panic("implement me")
}
