package ormutil

import "gorm.io/gorm"

// WhereWithParamMap using param
// paramMap is used to supply parameters to avoid sql injection
// refer to https://gorm.io/docs/advanced_query.html#Named-Argument
func WhereWithParamMap(db *gorm.DB, clause string, paramMap map[string]interface{}) *gorm.DB {
	// either clause or map is empty will cause error.
	if len(clause) == 0 {
		return db
	} else if len(paramMap) == 0 {
		return db.Where(clause)
	} else {
		return db.Where(clause, paramMap)
	}
}
