package query

import (
	"fmt"

	"gorm.io/gorm"
)

func HasPrefix(field string, value *string) func(db *gorm.DB) *gorm.DB {
	if value == nil {
		return self()
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s LIKE '%s%%'", field, *value))
	}
}

func HasSuffix(field string, value *string) func(db *gorm.DB) *gorm.DB {
	if value == nil {
		return self()
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s LIKE '%%%s'", field, *value))
	}
}
