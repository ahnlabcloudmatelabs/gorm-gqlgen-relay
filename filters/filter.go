package filters

import (
	"encoding/json"
)

type Filter[T any] struct {
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

func ParseFilterArray[T any](input any) ([]map[string]Filter[T], error) {
	if filters, ok := input.([]map[string]Filter[T]); ok {
		return filters, nil
	}

	data, _ := json.Marshal(input)

	var filters []map[string]Filter[T]
	if err := json.Unmarshal(data, &filters); err != nil {
		return nil, err
	}

	return filters, nil
}

func ParseFilterMap[T any](input any) (map[string]Filter[T], error) {
	if filters, ok := input.(map[string]Filter[T]); ok {
		return filters, nil
	}

	data, _ := json.Marshal(input)

	var filters map[string]Filter[T]
	if err := json.Unmarshal(data, &filters); err != nil {
		return nil, err
	}

	return filters, nil
}
