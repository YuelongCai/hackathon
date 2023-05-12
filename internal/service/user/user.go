package user

import (
	"gorm.io/gorm"
	"hackson/internal/model/rds/data"

	"hackson/internal/data/user"
	"hackson/internal/model/rds"
)

// Service for user
type Service struct {
	rds                  *rds.Rds
	badgetModel          *rds.BadgetModel
	userBadgetAssetModel *rds.UserBadgetAssetModel
}

// NewService .
func NewService(rds *rds.Rds, bm *rds.BadgetModel, ubam *rds.UserBadgetAssetModel) *Service {
	return &Service{rds: rds, badgetModel: bm, userBadgetAssetModel: ubam}
}

// HandleBehavior .
func (s *Service) HandleBehavior(ub *user.Behavior) error {
	badgets, err := s.badgetModel.FetchByTriggerEvent(ub.EventName)
	if err != nil {
		return err
	}
	for _, badget := range badgets {
		err = s.rds.Transaction(func(tx *gorm.DB) error {
			userBadgetAsset, err := s.userBadgetAssetModel.WithTx(tx).Get(ub.UserID, badget.ID)
			if err != nil {
				return err
			}
			if userBadgetAsset.CurrentState == -1 {
				// already finish all condition, return
				return nil
			} else if userBadgetAsset.CurrentState+ub.Count >= badget.Condition {
				userBadgetAsset.CurrentState = -1
			} else {
				// userBadgetAsset.CurrentState + ub.Count < badget.Condition
				userBadgetAsset.CurrentState += ub.Count
			}
			err = s.userBadgetAssetModel.WithTx(tx).Update(userBadgetAsset)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) ListBadgetAssets(userID int64, category string) ([]data.Badget, error) {
	badgetAssets, err := s.userBadgetAssetModel.GetBadgetAssetsByUser(userID, category)
	if err != nil {
		return nil, nil
	}
	badgets := make([]data.Badget, 0)
	for _, badgetAsset := range badgetAssets {
		badgets = append(badgets, badgetAsset.Badget)
	}
	return badgets, nil
}
