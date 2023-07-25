package filters

import (
	"github.com/cloudmatelabs/gorm-gqlgen-relay/query"
	"gorm.io/gorm"
)

type StringFilter struct {
	Equal              *string   `json:"equal,omitempty"`
	EqualFold          *string   `json:"equalFold,omitempty"`
	NotEqual           *string   `json:"notEqual,omitempty"`
	In                 *[]string `json:"in,omitempty"`
	NotIn              *[]string `json:"notIn,omitempty"`
	Contains           *string   `json:"contains,omitempty"`
	ContainsFold       *string   `json:"containsFold,omitempty"`
	GreaterThan        *string   `json:"gt,omitempty"`
	GreaterThanOrEqual *string   `json:"gte,omitempty"`
	LessThan           *string   `json:"lt,omitempty"`
	LessThanOrEqual    *string   `json:"lte,omitempty"`
	HasPrefix          *string   `json:"hasPrefix,omitempty"`
	HasSuffix          *string   `json:"hasSuffix,omitempty"`
	IsNull             *bool     `json:"isNull,omitempty"`
	IsNotNull          *bool     `json:"isNotNull,omitempty"`
}

func String(db *gorm.DB, field string, input interface{}) (*gorm.DB, error) {
	var filter Filter[string]
	if err := filter.Parse(input); err != nil {
		return db, err
	}

	db = db.Scopes(
		query.Equal(field, filter.Equal),
		query.NotEqual(field, filter.NotEqual),
		query.EqualFold(field, filter.EqualFold),
		query.In(field, filter.In),
		query.NotIn(field, filter.NotIn),
		query.Contains(field, filter.Contains),
		query.ContainsFold(field, filter.ContainsFold),
		query.GreaterThan(field, filter.GreaterThan),
		query.GreaterThanOrEqual(field, filter.GreaterThanOrEqual),
		query.LessThan(field, filter.LessThan),
		query.LessThanOrEqual(field, filter.LessThanOrEqual),
		query.HasPrefix(field, filter.HasPrefix),
		query.HasSuffix(field, filter.HasSuffix),
		query.IsNull(field, filter.IsNull),
		query.IsNotNull(field, filter.IsNotNull),
	)

	return db, nil
}
