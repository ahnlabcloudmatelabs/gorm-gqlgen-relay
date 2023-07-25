package filters

import (
	"github.com/cloudmatelabs/gorm-gqlgen-relay/query"
	"gorm.io/gorm"
)

type BoolFilter struct {
	Equal              *bool   `json:"equal,omitempty"`
	NotEqual           *bool   `json:"notEqual,omitempty"`
	In                 *[]bool `json:"in,omitempty"`
	NotIn              *[]bool `json:"notIn,omitempty"`
	GreaterThan        *bool   `json:"gt,omitempty"`
	GreaterThanOrEqual *bool   `json:"gte,omitempty"`
	LessThan           *bool   `json:"lt,omitempty"`
	LessThanOrEqual    *bool   `json:"lte,omitempty"`
	IsNull             *bool   `json:"isNull,omitempty"`
	IsNotNull          *bool   `json:"isNotNull,omitempty"`
}

func Bool(db *gorm.DB, field string, input interface{}) (*gorm.DB, error) {
	var filter Filter[bool]
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
