package data

import "time"

// UserBadgeAsset record of table `user_badge_asset`
type UserBadgeAsset struct {
	ID int64 `json:"id"`

	UserID       int64 `json:"user_id"`
	BadgeID     int64 `json:"badge_id"`
	CurrentState int   `json:"current_state"`

	Badge Badge `json:"badge,omitempty" gorm:"foreignKey:BadgeID"`

	CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName of `user_badge_asset`
func (r *UserBadgeAsset) TableName() string {
	return "user_badge_asset"
}
