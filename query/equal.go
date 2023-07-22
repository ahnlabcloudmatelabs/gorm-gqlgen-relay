package query

import (
	"fmt"

	"gorm.io/gorm"
)

func Equal[T any](field string, value *T) func(db *gorm.DB) *gorm.DB {
	if value == nil {
		return self()
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Where(field+" = ?", *value)
	}
}

func NotEqual[T any](field string, value *T) func(db *gorm.DB) *gorm.DB {
	if value == nil {
		return self()
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Where(field+" <> ?", *value)
	}
}

func EqualFold(field string, value *string) func(db *gorm.DB) *gorm.DB {
	if value == nil {
		return self()
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("LOWER(%s) = LOWER(?)", field), *value)
	}
}
