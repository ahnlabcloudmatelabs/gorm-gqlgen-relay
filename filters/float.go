package filters

import (
	"github.com/cloudmatelabs/gorm-gqlgen-relay/query"
)

type FloatFilter struct {
	Equal              *float64   `json:"equal,omitempty"`
	NotEqual           *float64   `json:"notEqual,omitempty"`
	In                 *[]float64 `json:"in,omitempty"`
	NotIn              *[]float64 `json:"notIn,omitempty"`
	GreaterThan        *float64   `json:"gt,omitempty"`
	GreaterThanOrEqual *float64   `json:"gte,omitempty"`
	LessThan           *float64   `json:"lt,omitempty"`
	LessThanOrEqual    *float64   `json:"lte,omitempty"`
	IsNull             *bool      `json:"isNull,omitempty"`
	IsNotNull          *bool      `json:"isNotNull,omitempty"`
}

func Float(field string, input any) (queryString string, values []any, err error) {
	var filter Filter[float64]
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
