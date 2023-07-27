package filters

import (
	"time"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/query"
)

type TimeFilter struct {
	Equal              *time.Time   `json:"equal,omitempty"`
	NotEqual           *time.Time   `json:"notEqual,omitempty"`
	In                 *[]time.Time `json:"in,omitempty"`
	NotIn              *[]time.Time `json:"notIn,omitempty"`
	GreaterThan        *time.Time   `json:"gt,omitempty"`
	GreaterThanOrEqual *time.Time   `json:"gte,omitempty"`
	LessThan           *time.Time   `json:"lt,omitempty"`
	LessThanOrEqual    *time.Time   `json:"lte,omitempty"`
	IsNull             *bool        `json:"isNull,omitempty"`
	IsNotNull          *bool        `json:"isNotNull,omitempty"`
}

func Time(field string, input any) (queryString string, values []any, err error) {
	var filter Filter[time.Time]
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
