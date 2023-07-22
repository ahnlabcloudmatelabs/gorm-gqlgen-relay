package query

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func Contains(field string, value *string) func(db *gorm.DB) *gorm.DB {
	if value == nil {
		return self()
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s LIKE '%%%s%%'", field, *value))
	}
}

func ContainsFold(field string, value *string) func(db *gorm.DB) *gorm.DB {
	if value == nil {
		return self()
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("LOWER(%s) LIKE '%%%s%%'", field, strings.ToLower(*value)))
	}
}
