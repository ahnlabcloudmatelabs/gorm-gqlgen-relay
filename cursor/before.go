package cursor

import (
	"fmt"

	"gorm.io/gorm"
)

func Before(db *gorm.DB, before *string, orderBy []map[string]string) (*gorm.DB, error) {
	if before == nil {
		return db, nil
	}

	cursor, err := decodeCursor(before)
	if err != nil {
		return nil, err
	}

	if len(orderBy) == 0 {
		return db.Where("id < ?", cursor[0]), nil
	}

	for i, order := range orderBy {
		if order["direction"] == "ASC" {
			db = db.Where(fmt.Sprintf("%s < ?", order["field"]), cursor[i])
		} else {
			db = db.Where(fmt.Sprintf("%s > ?", order["field"]), cursor[i])
		}
	}

	return db, nil
}
