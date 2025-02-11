package usecase

import (
	"context"
	"github.com/ardwiinoo/micro-music/musics/internal/domains/songs"
	"io"
	"net/http"
)

type StreamSongUseCase interface {
	Execute(ctx context.Context, songId string, rangeHeader string) (io.ReadCloser, int, string, error)
}

type streamSongUseCase struct {
	songRepository songs.SongRepository
}

func NewStreamSongUseCase(songRepository songs.SongRepository) StreamSongUseCase {
	return &streamSongUseCase{
		songRepository: songRepository,
	}
}

func (s streamSongUseCase) Execute(ctx context.Context, songId string, rangeHeader string) (io.ReadCloser, int, string, error) {
	song, err := s.songRepository.GetSongById(ctx, songId)
	if err != nil {
		return nil, 0, "", err
	}

	// Req file from firebase
	req, err := http.NewRequest("GET", song.Url, nil)
	if err != nil {
		return nil, 0, "", err
	}

	if rangeHeader != "" {
		req.Header.Set("Range", rangeHeader)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, "", err
	}

	return resp.Body, resp.StatusCode, resp.Header.Get("Content-Type"), nil
}
