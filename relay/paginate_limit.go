package relay

import "gorm.io/gorm"

func limit(db *gorm.DB, first, last *int) *gorm.DB {
	if first != nil {
		return db.Limit(*first)
	}

	if last != nil {
		return db.Limit(*last)
	}

	return db
}
