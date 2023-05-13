package badge

const (
	WatchedVideo = "watched_video"
	LoggedIn     = "logged_in"
)

func GetTriggerEvent() []string {
	return []string{WatchedVideo, LoggedIn}
}
