package translator

import "github.com/ardwiinoo/micro-music/musics/internal/commons/exceptions"

var DomainErrorMapping = map[string]exceptions.AppError{
	"DETAIL_SONG.NOT_CONTAIN_NEEDED_PROPERTY": exceptions.InvariantError("Detail song not contain needed property"),
	"ADD_SONG.NOT_CONTAIN_NEEDED_PROPERTY":    exceptions.InvariantError("Add song not contain needed property"),
}