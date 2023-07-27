package filters

import (
	"github.com/cloudmatelabs/gorm-gqlgen-relay/query"
)

type BooleanFilter struct {
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

func Boolean(field string, input any) (queryString string, values []any, err error) {
	var filter Filter[bool]
	if err = filter.Parse(input); err != nil {
		return
	}

	query.Equal(field, filter.Equal, &queryString, &values)
	query.NotEqual(field, filter.NotEqual, &queryString, &values)
	query.In(field, filter.In, &queryString, &values)
	query.NotIn(field, filter.NotIn, &queryString, &values)
	query.GreaterThan(field, filter.GreaterThan, &queryString, &values)
	query.GreaterThanOrEqual(field, filter.GreaterThanOrEqual, &queryString, &values)
	query.LessThan(field, filter.LessThan, &queryString, &values)
	query.LessThanOrEqual(field, filter.LessThanOrEqual, &queryString, &values)
	query.IsNull(field, filter.IsNull, &queryString)
	query.IsNotNull(field, filter.IsNotNull, &queryString)

	return
}
