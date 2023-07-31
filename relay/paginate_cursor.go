package relay

import (
	"github.com/cloudmatelabs/gorm-gqlgen-relay/cursor"
	"gorm.io/gorm"
)

func setAfter(db *gorm.DB, after *string, orderBy map[string]any, primaryKey string) (*gorm.DB, error) {
	return setCursor(db, after, "after", orderBy, primaryKey)
}

func setBefore(db *gorm.DB, before *string, orderBy map[string]any, primaryKey string) (*gorm.DB, error) {
	return setCursor(db, before, "before", orderBy, primaryKey)
}

func setCursor(db *gorm.DB, cur *string, direction string, orderBy map[string]any, primaryKey string) (*gorm.DB, error) {
	if cur == nil {
		return db, nil
	}

	var queries []string
	var args []any
	var err error

	if direction == "after" {
		queries, args, err = cursor.After(cur, orderBy, primaryKey)
	} else {
		queries, args, err = cursor.Before(cur, orderBy, primaryKey)
	}

	if err != nil {
		return db, err
	}

	for i := range queries {
		db = db.Where(queries[i], args[i])
	}

	return db, err
}
