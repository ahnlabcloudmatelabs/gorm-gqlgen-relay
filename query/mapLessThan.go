package query

import (
	"fmt"

	"gorm.io/gorm"
)

func MapLessThan(field, key string, value any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch db.Dialector.Name() {
		case "postgres":
			db = db.Where(fmt.Sprintf("(%s->%s)::jsonb < '?'", field, key), value)
		case "sqlserver":
			db = db.Where(fmt.Sprintf("JSON_VALUE(%s, '%s') < ?", field, key), value)
		default:
			db = db.Where(fmt.Sprintf("JSON_EXTRACT(%s, '%s') < ?", field, key), value)
		}
		return db
	}
}

func MapLessThanOrEqual(field, key string, value any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch db.Dialector.Name() {
		case "postgres":
			db = db.Where(fmt.Sprintf("(%s->%s)::jsonb <= '?'", field, key), value)
		case "sqlserver":
			db = db.Where(fmt.Sprintf("JSON_VALUE(%s, '%s') <= ?", field, key), value)
		default:
			db = db.Where(fmt.Sprintf("JSON_EXTRACT(%s, '%s') <= ?", field, key), value)
		}
		return db
	}
}
