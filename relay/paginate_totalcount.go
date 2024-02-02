package relay

import (
	"github.com/cloudmatelabs/gorm-gqlgen-relay/where"
	"gorm.io/gorm"
)

func getTotalCount[Model any](db *gorm.DB, w where.Where) (int64, error) {
	var totalCount int64
	var model Model
	err := db.Model(&model).
		Where(w.Query, w.Args...).
		Count(&totalCount).Error

	if err != nil {
		return 0, err
	}

	return totalCount, nil
}
