package badge

import (
	"hackson/internal/model/rds"
	"hackson/internal/model/rds/data"
)

// Service for badge
type Service struct {
	badgeModel *rds.BadgeModel
}

// NewService .
func NewService(bm *rds.BadgeModel) *Service {
	return &Service{badgeModel: bm}
}

// ListByFilterAndOrder list Badge by filters and orders
func (s *Service) ListByFilterAndOrder(filterClause string, paramMap map[string]interface{}, orderClause string, offset, limit int) ([]data.Badge, error) {
	return s.badgeModel.ListByFilterAndOrder(filterClause, paramMap, orderClause, offset, limit)
}

// CountByFilter get Badge count by filters
func (s *Service) CountByFilter(filterClause string, paramMap map[string]interface{}) (int64, error) {
	return s.badgeModel.CountByFilter(filterClause, paramMap)
}

// Register a new badge
func (s *Service) Register(badge *data.Badge) error {
	return s.badgeModel.Register(badge)
}

// Change a badge
func (s *Service) Change(badge *data.Badge) error {
	return s.badgeModel.Change(badge)
}
