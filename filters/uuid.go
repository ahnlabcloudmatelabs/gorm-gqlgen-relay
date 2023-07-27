package filters

import (
	"github.com/cloudmatelabs/gorm-gqlgen-relay/query"
	"github.com/google/uuid"
)

type UUIDFilter struct {
	Equal              *uuid.UUID   `json:"equal,omitempty"`
	NotEqual           *uuid.UUID   `json:"notEqual,omitempty"`
	In                 *[]uuid.UUID `json:"in,omitempty"`
	NotIn              *[]uuid.UUID `json:"notIn,omitempty"`
	GreaterThan        *uuid.UUID   `json:"gt,omitempty"`
	GreaterThanOrEqual *uuid.UUID   `json:"gte,omitempty"`
	LessThan           *uuid.UUID   `json:"lt,omitempty"`
	LessThanOrEqual    *uuid.UUID   `json:"lte,omitempty"`
	IsNull             *bool        `json:"isNull,omitempty"`
	IsNotNull          *bool        `json:"isNotNull,omitempty"`
}

func UUID(field string, input any) (queryString string, values []any, err error) {
	var filter Filter[uuid.UUID]
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
