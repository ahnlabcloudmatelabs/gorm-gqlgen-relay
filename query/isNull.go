package query

import (
	"gorm.io/gorm"
)

func IsNull(field string, value *bool) func(db *gorm.DB) *gorm.DB {
	if value == nil {
		return self()
	}

	query := field + " IS NULL"

	if !*value {
		query = field + " IS NOT NULL"
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Where(query)
	}
}

func IsNotNull(field string, value *bool) func(db *gorm.DB) *gorm.DB {
	if value == nil {
		return self()
	}

	query := field + " IS NULL"

	if !*value {
		query = field + " IS NOT NULL"
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Not(query)
	}
}
