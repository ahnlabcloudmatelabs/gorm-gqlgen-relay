package filters

import "time"

type StringFilter struct {
	Equal              *string  `json:"equal,omitempty"`
	EqualFold          *string  `json:"equalFold,omitempty"`
	NotEqual           *string  `json:"notEqual,omitempty"`
	In                 []string `json:"in,omitempty"`
	NotIn              []string `json:"notIn,omitempty"`
	Contains           *string  `json:"contains,omitempty"`
	ContainsFold       *string  `json:"containsFold,omitempty"`
	GreaterThan        *string  `json:"gt,omitempty"`
	GreaterThanOrEqual *string  `json:"gte,omitempty"`
	LessThan           *string  `json:"lt,omitempty"`
	LessThanOrEqual    *string  `json:"lte,omitempty"`
	HasPrefix          *string  `json:"hasPrefix,omitempty"`
	HasSuffix          *string  `json:"hasSuffix,omitempty"`
	IsNull             *bool    `json:"isNull,omitempty"`
}

type IntFilter struct {
	Equal              *int  `json:"equal,omitempty"`
	NotEqual           *int  `json:"notEqual,omitempty"`
	In                 []int `json:"in,omitempty"`
	NotIn              []int `json:"notIn,omitempty"`
	GreaterThan        *int  `json:"gt,omitempty"`
	GreaterThanOrEqual *int  `json:"gte,omitempty"`
	LessThan           *int  `json:"lt,omitempty"`
	LessThanOrEqual    *int  `json:"lte,omitempty"`
	IsNull             *bool `json:"isNull,omitempty"`
}

type FloatFilter struct {
	Equal              *float64  `json:"equal,omitempty"`
	NotEqual           *float64  `json:"notEqual,omitempty"`
	In                 []float64 `json:"in,omitempty"`
	NotIn              []float64 `json:"notIn,omitempty"`
	GreaterThan        *float64  `json:"gt,omitempty"`
	GreaterThanOrEqual *float64  `json:"gte,omitempty"`
	LessThan           *float64  `json:"lt,omitempty"`
	LessThanOrEqual    *float64  `json:"lte,omitempty"`
	IsNull             *bool     `json:"isNull,omitempty"`
}

type IDFilter struct {
	Equal              *string  `json:"equal,omitempty"`
	EqualFold          *string  `json:"equalFold,omitempty"`
	NotEqual           *string  `json:"notEqual,omitempty"`
	In                 []string `json:"in,omitempty"`
	NotIn              []string `json:"notIn,omitempty"`
	Contains           *string  `json:"contains,omitempty"`
	ContainsFold       *string  `json:"containsFold,omitempty"`
	GreaterThan        *string  `json:"gt,omitempty"`
	GreaterThanOrEqual *string  `json:"gte,omitempty"`
	LessThan           *string  `json:"lt,omitempty"`
	LessThanOrEqual    *string  `json:"lte,omitempty"`
	HasPrefix          *string  `json:"hasPrefix,omitempty"`
	HasSuffix          *string  `json:"hasSuffix,omitempty"`
	IsNull             *bool    `json:"isNull,omitempty"`
}

type TimeFilter struct {
	Equal              *time.Time   `json:"equal,omitempty"`
	NotEqual           *time.Time   `json:"notEqual,omitempty"`
	In                 []*time.Time `json:"in,omitempty"`
	NotIn              []*time.Time `json:"notIn,omitempty"`
	GreaterThan        *time.Time   `json:"gt,omitempty"`
	GreaterThanOrEqual *time.Time   `json:"gte,omitempty"`
	LessThan           *time.Time   `json:"lt,omitempty"`
	LessThanOrEqual    *time.Time   `json:"lte,omitempty"`
	IsNull             *bool        `json:"isNull,omitempty"`
}

type UUIDFilter struct {
	Equal              *string  `json:"equal,omitempty"`
	NotEqual           *string  `json:"notEqual,omitempty"`
	In                 []string `json:"in,omitempty"`
	NotIn              []string `json:"notIn,omitempty"`
	GreaterThan        *string  `json:"gt,omitempty"`
	GreaterThanOrEqual *string  `json:"gte,omitempty"`
	LessThan           *string  `json:"lt,omitempty"`
	LessThanOrEqual    *string  `json:"lte,omitempty"`
	IsNull             *bool    `json:"isNull,omitempty"`
}

type MapFilter struct {
	Equal              *MapEntry   `json:"equal,omitempty"`
	EqualFold          *MapEntry   `json:"equalFold,omitempty"`
	NotEqual           *MapEntry   `json:"notEqual,omitempty"`
	In                 *MapInEntry `json:"in,omitempty"`
	NotIn              *MapInEntry `json:"notIn,omitempty"`
	Contains           *MapEntry   `json:"contains,omitempty"`
	ContainsFold       *MapEntry   `json:"containsFold,omitempty"`
	GreaterThan        *MapEntry   `json:"gt,omitempty"`
	GreaterThanOrEqual *MapEntry   `json:"gte,omitempty"`
	LessThan           *MapEntry   `json:"lt,omitempty"`
	LessThanOrEqual    *MapEntry   `json:"lte,omitempty"`
	HasPrefix          *MapEntry   `json:"hasPrefix,omitempty"`
	HasSuffix          *MapEntry   `json:"hasSuffix,omitempty"`
	IsNull             *bool       `json:"isNull,omitempty"`
}

type MapEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MapInEntry struct {
	Key   string   `json:"key"`
	Value []string `json:"value,omitempty"`
}
