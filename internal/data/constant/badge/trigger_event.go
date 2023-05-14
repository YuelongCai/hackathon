package badge

const (
	Login            = "login"
	BirthdayLogin    = "birthday_login"
	AnniversaryLogin = "anniversary_login"
	WatchVideo       = "watch_video"
)

func GetTriggerEvent() []string {
	return []string{Login, BirthdayLogin, AnniversaryLogin, WatchVideo}
}
