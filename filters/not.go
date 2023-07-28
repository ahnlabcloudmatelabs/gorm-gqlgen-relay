package filters

import (
	"gorm.io/gorm"
)

func Not[T any](db *gorm.DB, input *T) (*gorm.DB, error) {
	if input == nil {
		return db, nil
	}

	filters, err := ParseFilterMap[any](*input)
	if err != nil {
		return db, err
	}

	db = db.Scopes(func(d *gorm.DB) *gorm.DB {
		for field, filter := range filters {
			queryString, values := createQuery(field, filter)
			if queryString == "" {
				continue
			}

			d = d.Not(queryString, values...)
		}
		return d
	})

	return db, err
}
