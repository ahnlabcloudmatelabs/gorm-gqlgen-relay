package filters

import (
	"encoding/json"
)

type Filter[T any] struct {
	Not                *Filter[T]   `json:"not,omitempty"`
	And                *[]Filter[T] `json:"and,omitempty"`
	Or                 *[]Filter[T] `json:"or,omitempty"`
	Equal              *T           `json:"equal,omitempty"`
	EqualFold          *T           `json:"equalFold,omitempty"`
	NotEqual           *T           `json:"notEqual,omitempty"`
	In                 *[]T         `json:"in,omitempty"`
	NotIn              *[]T         `json:"notIn,omitempty"`
	Contains           *T           `json:"contains,omitempty"`
	ContainsFold       *T           `json:"containsFold,omitempty"`
	GreaterThan        *T           `json:"gt,omitempty"`
	GreaterThanOrEqual *T           `json:"gte,omitempty"`
	LessThan           *T           `json:"lt,omitempty"`
	LessThanOrEqual    *T           `json:"lte,omitempty"`
	HasPrefix          *T           `json:"hasPrefix,omitempty"`
	HasSuffix          *T           `json:"hasSuffix,omitempty"`
	IsNull             *bool        `json:"isNull,omitempty"`
	IsNotNull          *bool        `json:"isNotNull,omitempty"`
}

func (filter *Filter[T]) Parse(input interface{}) error {
	_filter, ok := input.(Filter[T])
	if ok {
		filter = &_filter
		return nil
	}

	byteData, err := json.Marshal(input)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteData, &filter)
}
