package user

// Behavior record of user_behavior
type Behavior struct {
	EventName string `json:"event_name"`
	UserID    int64  `json:"user_id"`
	Count     int    `json:"count"`
}
