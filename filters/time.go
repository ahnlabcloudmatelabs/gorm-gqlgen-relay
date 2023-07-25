package filters

import (
	"time"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/query"
	"gorm.io/gorm"
)

type TimeFilter struct {
	Equal              *time.Time   `json:"equal,omitempty"`
	EqualFold          *time.Time   `json:"equalFold,omitempty"`
	NotEqual           *time.Time   `json:"notEqual,omitempty"`
	In                 *[]time.Time `json:"in,omitempty"`
	NotIn              *[]time.Time `json:"notIn,omitempty"`
	Contains           *time.Time   `json:"contains,omitempty"`
	ContainsFold       *time.Time   `json:"containsFold,omitempty"`
	GreaterThan        *time.Time   `json:"gt,omitempty"`
	GreaterThanOrEqual *time.Time   `json:"gte,omitempty"`
	LessThan           *time.Time   `json:"lt,omitempty"`
	LessThanOrEqual    *time.Time   `json:"lte,omitempty"`
	HasPrefix          *time.Time   `json:"hasPrefix,omitempty"`
	HasSuffix          *time.Time   `json:"hasSuffix,omitempty"`
	IsNull             *bool        `json:"isNull,omitempty"`
	IsNotNull          *bool        `json:"isNotNull,omitempty"`
}

func Time(db *gorm.DB, field string, input interface{}) (*gorm.DB, error) {
	var filter Filter[time.Time]
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
