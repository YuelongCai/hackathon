package badget

import (
	"hackson/internal/model/rds"
	"hackson/internal/model/rds/data"
)

// Service for badget
type Service struct {
	badgetModel *rds.BadgetModel
}

// NewService .
func NewService(bm *rds.BadgetModel) *Service {
	return &Service{badgetModel: bm}
}

// ListByFilterAndOrder list Badget by filters and orders
func (s *Service) ListByFilterAndOrder(filterClause string, paramMap map[string]interface{}, orderClause string, offset, limit int) ([]data.Badget, error) {
	return s.badgetModel.ListByFilterAndOrder(filterClause, paramMap, orderClause, offset, limit)
}

// CountByFilter get Badget count by filters
func (s *Service) CountByFilter(filterClause string, paramMap map[string]interface{}) (int64, error) {
	return s.badgetModel.CountByFilter(filterClause, paramMap)
}

// Register a new badget
func (s *Service) Register(badget *data.Badget) error {
	return s.badgetModel.Register(badget)
}

// Change a badget
func (s *Service) Change(badget *data.Badget) error {
	return s.badgetModel.Change(badget)
}
