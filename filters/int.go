package filters

import (
	"github.com/cloudmatelabs/gorm-gqlgen-relay/query"
	"gorm.io/gorm"
)

type IntFilter struct {
	Equal              *int   `json:"equal,omitempty"`
	NotEqual           *int   `json:"notEqual,omitempty"`
	In                 *[]int `json:"in,omitempty"`
	NotIn              *[]int `json:"notIn,omitempty"`
	GreaterThan        *int   `json:"gt,omitempty"`
	GreaterThanOrEqual *int   `json:"gte,omitempty"`
	LessThan           *int   `json:"lt,omitempty"`
	LessThanOrEqual    *int   `json:"lte,omitempty"`
	IsNull             *bool  `json:"isNull,omitempty"`
	IsNotNull          *bool  `json:"isNotNull,omitempty"`
}

func Int(db *gorm.DB, field string, input interface{}) (*gorm.DB, error) {
	var filter Filter[int]
	if err := filter.Parse(input); err != nil {
		return db, err
	}

	db = db.Scopes(
		query.Equal(field, filter.Equal),
		query.NotEqual(field, filter.NotEqual),
		query.In(field, filter.In),
		query.NotIn(field, filter.NotIn),
		query.GreaterThan(field, filter.GreaterThan),
		query.GreaterThanOrEqual(field, filter.GreaterThanOrEqual),
		query.LessThan(field, filter.LessThan),
		query.LessThanOrEqual(field, filter.LessThanOrEqual),
		query.IsNull(field, filter.IsNull),
		query.IsNotNull(field, filter.IsNotNull),
	)

	return db, nil
}
