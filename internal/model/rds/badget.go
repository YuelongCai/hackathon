package rds

import (
	"gorm.io/gorm"

	"hackson/internal/data/constant/badget"
	"hackson/internal/model/rds/data"
	"hackson/internal/util/ormutil"
)

// BadgetModel for table `badget`
type BadgetModel struct {
	*Rds
}

// NewBadgetModel .
func NewBadgetModel(rds *Rds) *BadgetModel {
	return &BadgetModel{rds}
}

// WithTx returns a shadow copy of BadgetModel with transaction
func (m *BadgetModel) WithTx(tx *gorm.DB) *BadgetModel {
	return &BadgetModel{&Rds{tx}}
}

// ListByFilterAndOrder list Badget by filters and orders
func (m *BadgetModel) ListByFilterAndOrder(filterClause string, paramMap map[string]interface{}, orderClause string, offset, limit int) ([]data.Badget, error) {
	var result []data.Badget
	res := ormutil.WhereWithParamMap(m.db.Model(data.Badget{}), filterClause, paramMap).
		Order(orderClause).
		Offset(offset)
	if limit > 0 {
		res = res.Limit(limit)
	}
	res = res.Find(&result)
	return result, res.Error
}

// CountByFilter get count of signals by filter
func (m *BadgetModel) CountByFilter(filterClause string, paramMap map[string]interface{}) (int64, error) {
	var count int64
	res := ormutil.WhereWithParamMap(m.db.Model(data.Badget{}), filterClause, paramMap).
		Count(&count)
	return count, res.Error
}

// Register a Badget
func (m *BadgetModel) Register(badget *data.Badget) error {
	err := m.db.Create(badget).Error
	return err
}

// Change a Badget
func (m *BadgetModel) Change(badget *data.Badget) error {
	err := m.db.Model(badget).Save(badget).Error
	return err
}

// FetchByTriggerEvent .
func (m *BadgetModel) FetchByTriggerEvent(triggerEvent string) ([]data.Badget, error) {
	var result []data.Badget
	err := m.db.Model(data.Badget{}).Find(&result, "trigger_event = ? and status = ?", triggerEvent, badget.Published).Error
	return result, err
}
