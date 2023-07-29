package paginate

import (
	"encoding/json"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func OrderBy(db *gorm.DB, orderBy []map[string]string, table *string, reverse bool) *gorm.DB {
	if len(orderBy) == 0 {
		return db
	}

	prefix := tablePrefix(table)

	for _, order := range orderBy {
		db = db.Order(
			fmt.Sprintf(
				"%s%s %s",
				prefix,
				order["field"],
				direction(order["direction"], reverse),
			),
		)
	}

	return db
}

func ParseOrderBy(orderBy any) ([]map[string]string, error) {
	filter := []map[string]string{}
	data, err := json.Marshal(orderBy)
	if err != nil {
		return filter, err
	}

	err = json.Unmarshal(data, &filter)

	return filter, err
}

func tablePrefix(table *string) string {
	if table == nil {
		return ""
	}

	return *table + "."
}

func direction(input string, reverse bool) string {
	if !reverse {
		return input
	}

	if strings.ToLower(input) == "asc" {
		return "DESC"
	}
	return "ASC"
}
