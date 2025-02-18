package usecase

import (
	"context"
	"encoding/json"
	"github.com/ardwiinoo/micro-music/musics/internal/commons/constants"
	"time"

	"github.com/ardwiinoo/micro-music/musics/internal/applications/cache"
	"github.com/ardwiinoo/micro-music/musics/internal/domains/songs"
	"github.com/ardwiinoo/micro-music/musics/internal/domains/songs/entities"
)

type GetListSongUseCase interface {
	Execute(ctx context.Context) (listSong []entities.DetailSong, isCached bool, err error)
}

type getListSongUseCase struct {
	songRepository songs.SongRepository
	redis          cache.CacheManager
}

func NewGetListSongUseCase(songRepository songs.SongRepository, redis cache.CacheManager) GetListSongUseCase {
	return &getListSongUseCase{
		songRepository: songRepository,
		redis:          redis,
	}
}

func (g *getListSongUseCase) Execute(ctx context.Context) (listSong []entities.DetailSong, isCached bool, err error) {
	cachedList, err := g.redis.Get(ctx, constants.RedisKey_ListSong)
	if err == nil {
		var cachedSongs []entities.DetailSong
		if err = json.Unmarshal([]byte(cachedList), &cachedSongs); err == nil {
			return cachedSongs, true, nil
		}
	}

	listSong, err = g.songRepository.GetListSong(ctx)
	if err != nil {
		return nil, false, err
	}

	data, err := json.Marshal(listSong)
	if err != nil {
		return listSong, false, err
	}

	if err = g.redis.Set(ctx, constants.RedisKey_ListSong, data, 30*time.Minute); err != nil {
		return listSong, false, err
	}

	return listSong, false, nil
}
