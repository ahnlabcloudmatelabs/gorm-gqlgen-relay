package query

import (
	"gorm.io/gorm"
)

func GreaterThan[T any](field string, value *T) func(db *gorm.DB) *gorm.DB {
	if value == nil {
		return self()
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Where(field+" > ?", *value)
	}
}

func GreaterThanOrEqual[T any](field string, value *T) func(db *gorm.DB) *gorm.DB {
	if value == nil {
		return self()
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Where(field+" >= ?", *value)
	}
}
