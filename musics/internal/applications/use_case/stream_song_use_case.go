package usecase

import (
	"context"
	"github.com/ardwiinoo/micro-music/musics/internal/domains/songs"
	"io"
	"net/http"
)

type StreamSongUseCase interface {
	Execute(ctx context.Context, songId string, rangeHeader string) (io.ReadCloser, int, string, http.Header, error)
}

type streamSongUseCase struct {
	songRepository songs.SongRepository
}

func NewStreamSongUseCase(songRepository songs.SongRepository) StreamSongUseCase {
	return &streamSongUseCase{
		songRepository: songRepository,
	}
}

func (s *streamSongUseCase) Execute(ctx context.Context, songID string, rangeHeader string) (io.ReadCloser, int, string, http.Header, error) {
	song, err := s.songRepository.GetSongById(ctx, songID)
	if err != nil {
		return nil, 0, "", nil, err
	}

	req, err := http.NewRequest("GET", song.Url, nil)
	if err != nil {
		return nil, 0, "", nil, err
	}

	if rangeHeader != "" {
		req.Header.Set("Range", rangeHeader)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, "", nil, err
	}

	return resp.Body,
		resp.StatusCode, // from firebase
		resp.Header.Get("Content-Type"),
		resp.Header,
		nil
}
