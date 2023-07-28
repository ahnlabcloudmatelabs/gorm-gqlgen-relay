package filters

import (
	"reflect"
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

	queryString, values = createQuery(field, filter)
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
