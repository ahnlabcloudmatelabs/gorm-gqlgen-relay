package filters

import (
	"fmt"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/query"
	"gorm.io/gorm"
)

type IDFilter[T any] struct {
	Not                *IDFilter[T]   `json:"not,omitempty"`
	And                *[]IDFilter[T] `json:"and,omitempty"`
	Or                 *[]IDFilter[T] `json:"or,omitempty"`
	Equal              *T             `json:"equal,omitempty"`
	EqualFold          *T             `json:"equalFold,omitempty"`
	NotEqual           *T             `json:"notEqual,omitempty"`
	In                 *[]T           `json:"in,omitempty"`
	NotIn              *[]T           `json:"notIn,omitempty"`
	Contains           *T             `json:"contains,omitempty"`
	ContainsFold       *T             `json:"containsFold,omitempty"`
	GreaterThan        *T             `json:"gt,omitempty"`
	GreaterThanOrEqual *T             `json:"gte,omitempty"`
	LessThan           *T             `json:"lt,omitempty"`
	LessThanOrEqual    *T             `json:"lte,omitempty"`
	HasPrefix          *T             `json:"hasPrefix,omitempty"`
	HasSuffix          *T             `json:"hasSuffix,omitempty"`
	IsNull             *bool          `json:"isNull,omitempty"`
	IsNotNull          *bool          `json:"isNotNull,omitempty"`
}

func ID[T any](db *gorm.DB, field string, input interface{}) (*gorm.DB, error) {
	var filter Filter[T]
	if err := filter.Parse(input); err != nil {
		return db, err
	}

	wheres := [](func(db *gorm.DB) *gorm.DB){
		query.Equal(field, filter.Equal),
		query.NotEqual(field, filter.NotEqual),
		query.In(field, filter.In),
		query.NotIn(field, filter.NotIn),
		query.GreaterThan(field, filter.GreaterThan),
		query.GreaterThanOrEqual(field, filter.GreaterThanOrEqual),
		query.LessThan(field, filter.LessThan),
		query.LessThanOrEqual(field, filter.LessThanOrEqual),
		query.IsNull(field, filter.IsNull),
		query.IsNotNull(field, filter.IsNotNull),
	}

	if fmt.Sprintf("%T", *new(T)) == "string" {
		toString := func(value *T) *string {
			if value == nil {
				return nil
			}

			_value := fmt.Sprintf("%v", *value)
			return &_value
		}

		wheres = append(wheres, []func(db *gorm.DB) *gorm.DB{
			query.EqualFold(field, toString(filter.EqualFold)),
			query.Contains(field, toString(filter.Contains)),
			query.ContainsFold(field, toString(filter.ContainsFold)),
			query.HasPrefix(field, toString(filter.HasPrefix)),
			query.HasSuffix(field, toString(filter.HasSuffix)),
		}...)
	}

	db = db.Scopes(wheres...)

	if filter.Not != nil {
		db = db.Scopes(func(d *gorm.DB) *gorm.DB {
			return d.Not(ID[T](d, field, *filter.Not))
		})
	}

	if filter.And != nil {
		for _, and := range *filter.And {
			_filter := and

			db = db.Scopes(func(d *gorm.DB) *gorm.DB {
				return d.Where(ID[T](d, field, _filter))
			})
		}
	}

	if filter.Or != nil {
		for _, or := range *filter.Or {
			_filter := or

			db = db.Scopes(func(d *gorm.DB) *gorm.DB {
				return d.Or(ID[T](d, field, _filter))
			})
		}
	}

	return db, nil
}
