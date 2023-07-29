package paginate

import "gorm.io/gorm"

func Limit(db *gorm.DB, first, last *int) *gorm.DB {
	if first != nil {
		db = db.Limit(*first)
	}

	if last != nil {
		db = db.Limit(*last)
	}

	return db
}
