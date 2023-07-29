package paginate

import "gorm.io/gorm"

func TotalCount[Model any](db *gorm.DB, model Model) (int64, error) {
	var count int64
	err := db.Model(model).Count(&count).Error

	return count, err
}
