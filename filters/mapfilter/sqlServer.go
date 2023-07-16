package mapfilter

import (
	"fmt"

	"gorm.io/gorm"
)

func MapEqualSqlServer(db *gorm.DB, field string, value map[string]interface{}, where bool) *gorm.DB {
	if where {
		for k, v := range value {
			db = db.Where(fmt.Sprintf("JSON_VALUE(%s, '$.%s') = ?", field, k), v)
		}
		return db
	}

	for k, v := range value {
		db = db.Or(fmt.Sprintf("JSON_VALUE(%s, '$.%s') = ?", field, k), v)
	}
	return db
}

func MapArrayEqualSqlServer(db *gorm.DB, field string, value []map[string]interface{}, where bool) *gorm.DB {
	for _, v := range value {
		db = MapEqualSqlServer(db, field, v, where)
	}
	return db
}

func MapNotEqualSqlServer(db *gorm.DB, field string, value map[string]interface{}, where bool) *gorm.DB {
	if where {
		for k, v := range value {
			db = db.Where(fmt.Sprintf("JSON_VALUE(%s, '$.%s') != ?", field, k), v)
		}
		return db
	}

	for k, v := range value {
		db = db.Or(fmt.Sprintf("JSON_VALUE(%s, '$.%s') != ?", field, k), v)
	}
	return db
}

func MapArrayNotEqualSqlServer(db *gorm.DB, field string, value []map[string]interface{}, where bool) *gorm.DB {
	for _, v := range value {
		db = MapNotEqualSqlServer(db, field, v, where)
	}
	return db
}

func MapInSqlServer(db *gorm.DB, field string, values []string, where bool) *gorm.DB {
	if where {
		for _, value := range values {
			db = db.Where(fmt.Sprintf("JSON_VALUE(%s, '$.%s') IS NOT NULL", field, value))
		}
		return db
	}

	for _, value := range values {
		db = db.Or(fmt.Sprintf("JSON_VALUE(%s, '$.%s') IS NOT NULL", field, value))
	}
	return db
}

func MapNotInSqlServer(db *gorm.DB, field string, values []string, where bool) *gorm.DB {
	if where {
		for _, value := range values {
			db = db.Where(fmt.Sprintf("JSON_VALUE(%s, '$.%s') IS NULL", field, value))
		}
		return db
	}

	for _, value := range values {
		db = db.Or(fmt.Sprintf("JSON_VALUE(%s, '$.%s') IS NULL", field, value))
	}
	return db
}
