package data

import "time"

// Badge record of table `badge`
type Badge struct {
	ID int64 `json:"id,omitempty"`

	Name         string `json:"name"`
	Image        string `json:"image"`
	Category     string `json:"category"`
	Description  string `json:"description"`
	Rarity       string `json:"rarity"`
	TriggerEvent string `json:"trigger_event"`
	Condition    int    `json:"condition"`
	Status       string `json:"status"`

	CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName of `badge`
func (r *Badge) TableName() string {
	return "badge"
}
