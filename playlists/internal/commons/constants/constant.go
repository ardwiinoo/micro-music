package constants

type ContextKey string

var (
	PublicIDContextKey ContextKey = "public_id"
)

var QueueTypes = struct {
	EmailQueue string
}{
	EmailQueue: "email_events",
}

var EventType = struct {
	UserRegistered string
	ExportPlaylist string
}{
	UserRegistered: "user_registered",
	ExportPlaylist: "export_playlist",
}

const (
	UserRoleID  int = 1
	AdminRoleID int = 2
)
