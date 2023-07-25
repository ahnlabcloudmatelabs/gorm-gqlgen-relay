package query

import (
	"fmt"

	"gorm.io/gorm"
)

func MapIn(field, key string, value []any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch db.Dialector.Name() {
		case "postgres":
			db = db.Where(fmt.Sprintf("(%s->%s)::jsonb IN '?'", field, key), value)
		case "sqlserver":
			db = db.Where(fmt.Sprintf("JSON_VALUE(%s, '%s') IN ?", field, key), value)
		default:
			db = db.Where(fmt.Sprintf("JSON_EXTRACT(%s, '%s') IN ?", field, key), value)
		}
		return db
	}
}

func MapNotIn(field, key string, value []any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch db.Dialector.Name() {
		case "postgres":
			db = db.Where(fmt.Sprintf("(%s->%s)::jsonb NOT IN '?'", field, key), value)
		case "sqlserver":
			db = db.Where(fmt.Sprintf("JSON_VALUE(%s, '%s') NOT IN ?", field, key), value)
		default:
			db = db.Where(fmt.Sprintf("JSON_EXTRACT(%s, '%s') NOT IN ?", field, key), value)
		}
		return db
	}
}
