package query

import (
	"gorm.io/gorm"
)

func self() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

func isMSSQL(db *gorm.DB) bool {
	return db.Dialector.Name() == "sqlserver"
}
