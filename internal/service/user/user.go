package user

import (
	"gorm.io/gorm"
	"hackathon/internal/model/rds/data"

	"hackathon/internal/data/user"
	"hackathon/internal/model/rds"
)

// Service for user
type Service struct {
	rds                 *rds.Rds
	badgeModel          *rds.BadgeModel
	userBadgeAssetModel *rds.UserBadgeAssetModel
}

// NewService .
func NewService(rds *rds.Rds, bm *rds.BadgeModel, ubam *rds.UserBadgeAssetModel) *Service {
	return &Service{rds: rds, badgeModel: bm, userBadgeAssetModel: ubam}
}

// HandleBehavior .
func (s *Service) HandleBehavior(ub *user.Behavior) error {
	badges, err := s.badgeModel.FetchByTriggerEvent(ub.EventName)
	if err != nil {
		return err
	}
	for _, badge := range badges {
		err = s.rds.Transaction(func(tx *gorm.DB) error {
			userBadgeAsset, err := s.userBadgeAssetModel.WithTx(tx).Get(ub.UserID, badge.ID)
			if err != nil {
				return err
			}
			if userBadgeAsset.CurrentState == -1 {
				// already finish all condition, return
				return nil
			} else if userBadgeAsset.CurrentState+ub.Count >= badge.Condition {
				userBadgeAsset.CurrentState = -1
			} else {
				// userBadgeAsset.CurrentState + ub.Count < badge.Condition
				userBadgeAsset.CurrentState += ub.Count
			}
			err = s.userBadgeAssetModel.WithTx(tx).Update(userBadgeAsset)
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

func (s *Service) ListBadgeAssets(userID int64, category string) ([]data.Badge, error) {
	badgeAssets, err := s.userBadgeAssetModel.GetBadgeAssetsByUser(userID, category)
	if err != nil {
		return nil, nil
	}
	badges := make([]data.Badge, 0)
	for _, badgeAsset := range badgeAssets {
		badges = append(badges, badgeAsset.Badge)
	}
	return badges, nil
}
