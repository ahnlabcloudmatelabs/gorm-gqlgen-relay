package filters

import (
	"gorm.io/gorm"
)

func Where[T any](db *gorm.DB, input *T) (*gorm.DB, error) {
	if input == nil {
		return db, nil
	}

	filters, err := ParseFilterMap[any](*input)
	if err != nil {
		return db, err
	}

	for field, filter := range filters {
		queryString, values := createQuery(field, filter)
		db = db.Where(queryString, values...)
	}

	return db, err
}
