package query

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func MapEqual(field string, key string, value any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch db.Dialector.Name() {
		case "postgres":
			db = db.Where(fmt.Sprintf("(%s->%s)::jsonb = '?'", field, key), value)
		case "sqlserver":
			db = db.Where(fmt.Sprintf("JSON_VALUE(%s, '%s') = ?", field, key), value)
		default:
			db = db.Where(fmt.Sprintf("JSON_EXTRACT(%s, '%s') = ?", field, key), value)
		}

		return db
	}
}

func MapNotEqual(field string, key string, value any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch db.Dialector.Name() {
		case "postgres":
			db = db.Where(fmt.Sprintf("(%s->%s)::jsonb != '?'", field, key), value)
		case "sqlserver":
			db = db.Where(fmt.Sprintf("JSON_VALUE(%s, '%s') != ?", field, key), value)
		default:
			db = db.Where(fmt.Sprintf("JSON_EXTRACT(%s, '%s') != ?", field, key), value)
		}

		return db
	}
}

func MapEqualFold(field string, key string, value string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch db.Dialector.Name() {
		case "postgres":
			db = db.Where(
				fmt.Sprintf("LOWER((%s->%s)::text) = '?'", field, key),
				strings.ToLower(value),
			)
		case "sqlserver":
			db = db.Where(
				fmt.Sprintf("LOWER(JSON_VALUE(%s, '%s')) = ?", field, key),
				strings.ToLower(value),
			)
		default:
			db = db.Where(
				fmt.Sprintf("LOWER(JSON_EXTRACT(%s, '%s')) = ?", field, key),
				strings.ToLower(value),
			)
		}

		return db
	}
}
