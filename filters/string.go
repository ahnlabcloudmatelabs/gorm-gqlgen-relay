package filters

import (
	"github.com/cloudmatelabs/gorm-gqlgen-relay/query"
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

func String(field string, input any) (queryString string, values []any, err error) {
	var filter Filter[string]
	if err = filter.Parse(input); err != nil {
		return
	}

	query.Equal(field, filter.Equal, &queryString, &values)
	query.EqualFold(field, filter.EqualFold, &queryString, &values)
	query.NotEqual(field, filter.NotEqual, &queryString, &values)
	query.In(field, filter.In, &queryString, &values)
	query.NotIn(field, filter.NotIn, &queryString, &values)
	query.Contains(field, filter.Contains, &queryString, &values)
	query.ContainsFold(field, filter.ContainsFold, &queryString, &values)
	query.GreaterThan(field, filter.GreaterThan, &queryString, &values)
	query.GreaterThanOrEqual(field, filter.GreaterThanOrEqual, &queryString, &values)
	query.LessThan(field, filter.LessThan, &queryString, &values)
	query.LessThanOrEqual(field, filter.LessThanOrEqual, &queryString, &values)
	query.HasPrefix(field, filter.HasPrefix, &queryString, &values)
	query.HasSuffix(field, filter.HasSuffix, &queryString, &values)
	query.IsNull(field, filter.IsNull, &queryString)
	query.IsNotNull(field, filter.IsNotNull, &queryString)

	return
}
