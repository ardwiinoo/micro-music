package constants

type ContextKey string

var (
	PublicIDContextKey ContextKey = "public_id"
)

const (
	UserRoleID  int = 1
	AdminRoleID int = 2
)

const (
	RedisKey_ListSong string = "songs:list"
)
