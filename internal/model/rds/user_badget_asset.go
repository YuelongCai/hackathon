package rds

import (
	"gorm.io/gorm"
	"hackson/internal/data/constant/badget"

	"hackson/internal/model/rds/data"
)

// UserBadgetAssetModel for table `user_badget_asset`
type UserBadgetAssetModel struct {
	*Rds
}

// NewUserBadgetAssetModel .
func NewUserBadgetAssetModel(rds *Rds) *UserBadgetAssetModel {
	return &UserBadgetAssetModel{rds}
}

// WithTx returns a shadow copy of UserBadgetAssetModel with transaction
func (m *UserBadgetAssetModel) WithTx(tx *gorm.DB) *UserBadgetAssetModel {
	return &UserBadgetAssetModel{&Rds{tx}}
}

// Get one record
func (m *UserBadgetAssetModel) Get(userID, badgetID int64) (*data.UserBadgetAsset, error) {
	userBadgetAsset := data.UserBadgetAsset{UserID: userID, BadgetID: badgetID}
	err := m.db.Model(data.UserBadgetAsset{}).FirstOrCreate(&userBadgetAsset).Error
	return &userBadgetAsset, err
}

// Update one record
func (m *UserBadgetAssetModel) Update(userBadgetAsset *data.UserBadgetAsset) error {
	err := m.db.Model(userBadgetAsset).Save(userBadgetAsset).Error
	return err
}

// GetBadgetsByUser .
func (m *UserBadgetAssetModel) GetBadgetAssetsByUser(userID int64) ([]data.UserBadgetAsset, error) {
	var userBadgetAssets []data.UserBadgetAsset
	err := m.db.Model(data.UserBadgetAsset{UserID: userID}).Preload("Badget", "Status = ?", badget.Published).Find(&userBadgetAssets).Error
	return userBadgetAssets, err
}
