package cursor

import (
	"fmt"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/order"
	"gorm.io/gorm"
)

func Before[T any](db *gorm.DB, before *string, orderBy []*T) (*gorm.DB, error) {
	if before == nil {
		return db, nil
	}

	_orderBy := order.ParseOrderBy(orderBy)
	cursor, err := decodeCursor(before)
	if err != nil {
		return nil, err
	}

	if len(_orderBy) == 0 {
		return db.Where("id < ?", cursor[0]), nil
	}

	for i, order := range _orderBy {
		if order["direction"] == "ASC" {
			db = db.Where(fmt.Sprintf("%s < ?", order["field"]), cursor[i])
		} else {
			db = db.Where(fmt.Sprintf("%s > ?", order["field"]), cursor[i])
		}
	}

	return db, nil
}
