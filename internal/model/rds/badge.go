package rds

import (
	"gorm.io/gorm"

	"hackson/internal/data/constant/badge"
	"hackson/internal/model/rds/data"
	"hackson/internal/util/ormutil"
)

// BadgeModel for table `badge`
type BadgeModel struct {
	*Rds
}

// NewBadgeModel .
func NewBadgeModel(rds *Rds) *BadgeModel {
	return &BadgeModel{rds}
}

// WithTx returns a shadow copy of BadgeModel with transaction
func (m *BadgeModel) WithTx(tx *gorm.DB) *BadgeModel {
	return &BadgeModel{&Rds{tx}}
}

// ListByFilterAndOrder list Badge by filters and orders
func (m *BadgeModel) ListByFilterAndOrder(filterClause string, paramMap map[string]interface{}, orderClause string, offset, limit int) ([]data.Badge, error) {
	var result []data.Badge
	res := ormutil.WhereWithParamMap(m.db.Model(data.Badge{}), filterClause, paramMap).
		Order(orderClause).
		Offset(offset)
	if limit > 0 {
		res = res.Limit(limit)
	}
	res = res.Find(&result)
	return result, res.Error
}

// CountByFilter get count of signals by filter
func (m *BadgeModel) CountByFilter(filterClause string, paramMap map[string]interface{}) (int64, error) {
	var count int64
	res := ormutil.WhereWithParamMap(m.db.Model(data.Badge{}), filterClause, paramMap).
		Count(&count)
	return count, res.Error
}

// Register a Badge
func (m *BadgeModel) Register(badge *data.Badge) error {
	err := m.db.Create(badge).Error
	return err
}

// Change a Badge
func (m *BadgeModel) Change(badge *data.Badge) error {
	err := m.db.Model(badge).Save(badge).Error
	return err
}

// FetchByTriggerEvent .
func (m *BadgeModel) FetchByTriggerEvent(triggerEvent string) ([]data.Badge, error) {
	var result []data.Badge
	err := m.db.Model(data.Badge{}).Find(&result, "trigger_event = ? and status = ?", triggerEvent, badge.Published).Error
	return result, err
}
