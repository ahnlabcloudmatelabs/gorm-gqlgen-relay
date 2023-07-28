package count

import "gorm.io/gorm"

func Count(db *gorm.DB, model any) int64 {
	var count int64
	db.Model(model).Count(&count)

	return count
}
