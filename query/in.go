package query

import (
	"gorm.io/gorm"
)

func In[T any](field string, value *T) func(db *gorm.DB) *gorm.DB {
	if value == nil {
		return self()
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Where(field+" IN ?", *value)
	}
}

func NotIn[T any](field string, value *T) func(db *gorm.DB) *gorm.DB {
	if value == nil {
		return self()
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Where(field+" NOT IN ?", *value)
	}
}
