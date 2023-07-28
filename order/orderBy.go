package order

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

func OrderBy(db *gorm.DB, orderBy []map[string]string, tableName string) *gorm.DB {
	if len(orderBy) == 0 {
		return db
	}

	prefix := tablePrefix(tableName)

	for _, order := range orderBy {
		db = db.Order(fmt.Sprintf("%s%s %s", prefix, order["field"], order["direction"]))
	}

	return db
}

func ParseOrderBy(orderBy any) []map[string]string {
	filter := []map[string]string{}
	data, _ := json.Marshal(orderBy)
	json.Unmarshal(data, &filter)

	return filter
}

func tablePrefix(tableName string) string {
	if tableName == "" {
		return ""
	}

	return tableName + "."
}
