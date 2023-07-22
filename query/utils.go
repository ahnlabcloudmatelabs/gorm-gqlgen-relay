package query

import "gorm.io/gorm"

func operator(positive, negative string, not bool) string {
	if not {
		return negative
	}

	return positive
}

func clause(db *gorm.DB, or bool) func(query interface{}, args ...interface{}) *gorm.DB {
	if or {
		return db.Or
	}

	return db.Where
}

func boolPointer(b bool) *bool {
	return &b
}

func isArray(value interface{}) bool {
	_, ok := value.([]interface{})
	return ok
}

func isMap(value interface{}) bool {
	_, ok := value.(map[string]interface{})
	return ok
}

func self() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

func isMSSQL(db *gorm.DB) bool {
	return db.Dialector.Name() == "sqlserver"
}
