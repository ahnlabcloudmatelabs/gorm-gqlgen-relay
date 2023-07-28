package filters

import (
	"gorm.io/gorm"
)

func And(db *gorm.DB, input any) (*gorm.DB, error) {
	if input == nil {
		return db, nil
	}

	filters, err := ParseFilterArray[any](input)
	if err != nil {
		return db, err
	}

	db = db.Scopes(func(d *gorm.DB) *gorm.DB {
		for _, filter := range filters {
			for field, f := range filter {
				queryString, values := createQuery(field, f)
				if queryString == "" {
					continue
				}

				d = d.Where(queryString, values...)
			}
		}
		return d
	})

	return db, err
}
