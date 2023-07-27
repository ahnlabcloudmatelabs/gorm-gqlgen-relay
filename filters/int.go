package filters

import (
	"github.com/cloudmatelabs/gorm-gqlgen-relay/query"
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

func Int(field string, input any) (queryString string, values []any, err error) {
	var filter Filter[int]
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
