package rds

import (
	"gorm.io/gorm"
	"hackson/internal/data/constant/badge"

	"hackson/internal/model/rds/data"
)

// UserBadgeAssetModel for table `user_badge_asset`
type UserBadgeAssetModel struct {
	*Rds
}

// NewUserBadgeAssetModel .
func NewUserBadgeAssetModel(rds *Rds) *UserBadgeAssetModel {
	return &UserBadgeAssetModel{rds}
}

// WithTx returns a shadow copy of UserBadgeAssetModel with transaction
func (m *UserBadgeAssetModel) WithTx(tx *gorm.DB) *UserBadgeAssetModel {
	return &UserBadgeAssetModel{&Rds{tx}}
}

// Get one record
func (m *UserBadgeAssetModel) Get(userID, badgeID int64) (*data.UserBadgeAsset, error) {
	userBadgeAsset := data.UserBadgeAsset{UserID: userID, BadgeID: badgeID}
	err := m.db.Model(data.UserBadgeAsset{}).FirstOrCreate(&userBadgeAsset).Error
	return &userBadgeAsset, err
}

// Update one record
func (m *UserBadgeAssetModel) Update(userBadgeAsset *data.UserBadgeAsset) error {
	err := m.db.Model(userBadgeAsset).Save(userBadgeAsset).Error
	return err
}

// GetBadgesByUser .
func (m *UserBadgeAssetModel) GetBadgeAssetsByUser(userID int64, category string) ([]data.UserBadgeAsset, error) {
	var userBadgeAssets []data.UserBadgeAsset
	db := m.db.Model(data.UserBadgeAsset{UserID: userID})
	if category != "" {
		db = db.Debug().InnerJoins("Badge").Where("Badge.Status = ? and Badge.Category = ?", badge.Published, category)
	} else {
		db = db.Debug().InnerJoins("Badge").Where("Badge.Status = ?", badge.Published)
	}
	err := db.Where("current_state = -1").Find(&userBadgeAssets).Error
	return userBadgeAssets, err
}
