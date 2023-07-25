package filters

import (
	"encoding/json"
	"fmt"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/query"
	"gorm.io/gorm"
)

type MapFilter struct {
	Equal              *MapOneFilter  `json:"equal,omitempty"`
	EqualFold          *MapOneFilter  `json:"equalFold,omitempty"`
	NotEqual           *MapOneFilter  `json:"notEqual,omitempty"`
	In                 *MapManyFilter `json:"in,omitempty"`
	NotIn              *MapManyFilter `json:"notIn,omitempty"`
	Contains           *MapOneFilter  `json:"contains,omitempty"`
	ContainsFold       *MapOneFilter  `json:"containsFold,omitempty"`
	GreaterThan        *MapOneFilter  `json:"gt,omitempty"`
	GreaterThanOrEqual *MapOneFilter  `json:"gte,omitempty"`
	LessThan           *MapOneFilter  `json:"lt,omitempty"`
	LessThanOrEqual    *MapOneFilter  `json:"lte,omitempty"`
	HasPrefix          *MapOneFilter  `json:"hasPrefix,omitempty"`
	HasSuffix          *MapOneFilter  `json:"hasSuffix,omitempty"`
	Exists             *string        `json:"exists,omitempty"`
	NotExists          *string        `json:"notExists,omitempty"`
	IsNull             *bool          `json:"isNull,omitempty"`
	IsNotNull          *bool          `json:"isNotNull,omitempty"`
}

type MapOneFilter struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

type MapManyFilter struct {
	Key    string `json:"key"`
	Values []any  `json:"values"`
}

func Map(db *gorm.DB, field string, input interface{}) (*gorm.DB, error) {
	filter, err := parseMapWhere(input)
	if err != nil {
		return db, err
	}

	scopes := []func(db *gorm.DB) *gorm.DB{
		query.IsNull(field, filter.IsNull),
		query.IsNotNull(field, filter.IsNotNull),
	}

	if filter.Equal != nil {
		scopes = append(scopes, query.MapEqual(field, filter.Equal.Key, filter.Equal.Value))
	}

	if filter.NotEqual != nil {
		scopes = append(scopes, query.MapNotEqual(field, filter.NotEqual.Key, filter.NotEqual.Value))
	}

	if filter.EqualFold != nil {
		if _, ok := filter.EqualFold.Value.(string); !ok {
			return db, fmt.Errorf("EqualFold.Value must be string")
		}

		scopes = append(scopes, query.MapEqualFold(field, filter.EqualFold.Key, filter.EqualFold.Value.(string)))
	}

	if filter.In != nil {
		scopes = append(scopes, query.MapIn(field, filter.In.Key, filter.In.Values))
	}

	if filter.NotIn != nil {
		scopes = append(scopes, query.MapIn(field, filter.NotIn.Key, filter.NotIn.Values))
	}

	if filter.GreaterThan != nil {
		scopes = append(scopes, query.MapGreaterThan(field, filter.GreaterThan.Key, filter.GreaterThan.Value))
	}

	if filter.GreaterThanOrEqual != nil {
		scopes = append(scopes, query.MapGreaterThanOrEqual(field, filter.GreaterThanOrEqual.Key, filter.GreaterThanOrEqual.Value))
	}

	if filter.LessThan != nil {
		scopes = append(scopes, query.MapLessThan(field, filter.LessThan.Key, filter.LessThan.Value))
	}

	if filter.LessThanOrEqual != nil {
		scopes = append(scopes, query.MapLessThanOrEqual(field, filter.LessThanOrEqual.Key, filter.LessThanOrEqual.Value))
	}

	db = db.Scopes(scopes...)

	return db, nil
}

func parseMapWhere(input any) (filter MapFilter, err error) {
	_filter, ok := input.(MapFilter)
	if ok {
		filter = _filter
		return
	}

	var byteData []byte
	byteData, err = json.Marshal(input)
	if err != nil {
		return
	}

	err = json.Unmarshal(byteData, &filter)
	return
}
