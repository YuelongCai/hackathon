package data

import "time"

// UserBadgetAsset record of table `user_badget_asset`
type UserBadgetAsset struct {
	ID int64 `json:"id"`

	UserID       int64 `json:"user_id"`
	BadgetID     int64 `json:"badget_id"`
	CurrentState int   `json:"current_state"`

	Badget Badget `json:"badget,omitempty" gorm:"foreignKey:BadgetID"`

	CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName of `user_badget_asset`
func (r *UserBadgetAsset) TableName() string {
	return "user_badget_asset"
}
