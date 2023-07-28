package order

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

func OrderBy(db *gorm.DB, orderBy any, tableName string) *gorm.DB {
	filter := ParseOrderBy(orderBy)
	if len(filter) == 0 {
		return db
	}

	prefix := tablePrefix(tableName)

	for _, v := range filter {
		db = db.Order(fmt.Sprintf("%s%s %s", prefix, v["field"], v["direction"]))
	}

	return db
}

func ParseOrderBy(orderBy any) []map[string]any {
	filter := []map[string]any{}
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
