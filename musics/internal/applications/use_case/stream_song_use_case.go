package usecase

import (
	"context"
	"fmt"
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

func (s *streamSongUseCase) Execute(ctx context.Context, songID string, rangeHeader string) (io.ReadCloser, int, string, error) {
	song, err := s.songRepository.GetSongById(ctx, songID)
	if err != nil {
		return nil, 0, "", err
	}

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

	contentLength := resp.ContentLength
	contentRange := resp.Header.Get("Content-Range")

	statusCode := http.StatusOK
	if rangeHeader != "" {
		statusCode = http.StatusPartialContent
	}

	fmt.Println("Status Code:", statusCode)
	fmt.Println("Content-Type:", resp.Header.Get("Content-Type"))
	fmt.Println("Content-Length:", contentLength)
	fmt.Println("Content-Range:", contentRange)

	return resp.Body, statusCode, resp.Header.Get("Content-Type"), nil
}
