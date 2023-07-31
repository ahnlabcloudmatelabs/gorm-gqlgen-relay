package relay

import "gorm.io/gorm"

func getTotalCount[Model any](db *gorm.DB) (int64, error) {
	var totalCount int64
	var model Model
	if err := db.Model(&model).Count(&totalCount).Error; err != nil {
		return 0, err
	}
	return totalCount, nil
}
