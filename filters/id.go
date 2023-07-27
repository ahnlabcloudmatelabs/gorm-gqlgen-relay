package filters

import (
	"reflect"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/query"
)

type IDFilter[T any] struct {
	Equal              *T    `json:"equal,omitempty"`
	EqualFold          *T    `json:"equalFold,omitempty"`
	NotEqual           *T    `json:"notEqual,omitempty"`
	In                 *[]T  `json:"in,omitempty"`
	NotIn              *[]T  `json:"notIn,omitempty"`
	Contains           *T    `json:"contains,omitempty"`
	ContainsFold       *T    `json:"containsFold,omitempty"`
	GreaterThan        *T    `json:"gt,omitempty"`
	GreaterThanOrEqual *T    `json:"gte,omitempty"`
	LessThan           *T    `json:"lt,omitempty"`
	LessThanOrEqual    *T    `json:"lte,omitempty"`
	HasPrefix          *T    `json:"hasPrefix,omitempty"`
	HasSuffix          *T    `json:"hasSuffix,omitempty"`
	IsNull             *bool `json:"isNull,omitempty"`
	IsNotNull          *bool `json:"isNotNull,omitempty"`
}

func ID(field string, input any) (queryString string, values []any, err error) {
	var filter Filter[any]
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

	appendStringIDQuery(field, filter.EqualFold, &queryString, &values, query.EqualFold)
	appendStringIDQuery(field, filter.Contains, &queryString, &values, query.Contains)
	appendStringIDQuery(field, filter.ContainsFold, &queryString, &values, query.ContainsFold)
	appendStringIDQuery(field, filter.HasPrefix, &queryString, &values, query.HasPrefix)
	appendStringIDQuery(field, filter.HasSuffix, &queryString, &values, query.HasSuffix)

	return
}

func appendStringIDQuery(field string, input *any, queryString *string, values *[]any, callback func(string, *string, *string, *[]any)) {
	if input == nil {
		return
	}

	if reflect.ValueOf(*input).Kind() != reflect.String {
		return
	}

	value := (*input).(string)
	callback(field, &value, queryString, values)
}
