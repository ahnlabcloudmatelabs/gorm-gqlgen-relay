package filters

import (
	"encoding/json"

	"gorm.io/gorm"
)

func Do(db *gorm.DB, input any) *gorm.DB {
	if input == nil {
		return db
	}

	filter := parseFilter(input)

	db, _ = Where(db, &filter)
	db, _ = And(db, filter["and"])
	db, _ = Or(db, filter["or"])
	db, _ = Not(db, filter["not"])
	return db
}

func parseFilter(input any) (filter map[string]any) {
	if input == nil {
		return nil
	}

	data, _ := json.Marshal(input)
	json.Unmarshal(data, &filter)
	return
}
