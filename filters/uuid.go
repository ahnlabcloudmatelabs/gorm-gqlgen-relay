package filters

import (
	"github.com/cloudmatelabs/gorm-gqlgen-relay/query"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UUIDFilter struct {
	Not                *UUIDFilter   `json:"not,omitempty"`
	And                *[]UUIDFilter `json:"and,omitempty"`
	Or                 *[]UUIDFilter `json:"or,omitempty"`
	Equal              *uuid.UUID    `json:"equal,omitempty"`
	NotEqual           *uuid.UUID    `json:"notEqual,omitempty"`
	In                 *[]uuid.UUID  `json:"in,omitempty"`
	NotIn              *[]uuid.UUID  `json:"notIn,omitempty"`
	GreaterThan        *uuid.UUID    `json:"gt,omitempty"`
	GreaterThanOrEqual *uuid.UUID    `json:"gte,omitempty"`
	LessThan           *uuid.UUID    `json:"lt,omitempty"`
	LessThanOrEqual    *uuid.UUID    `json:"lte,omitempty"`
	IsNull             *bool         `json:"isNull,omitempty"`
	IsNotNull          *bool         `json:"isNotNull,omitempty"`
}

func UUID(db *gorm.DB, field string, input interface{}) (*gorm.DB, error) {
	var filter Filter[uuid.UUID]
	if err := filter.Parse(input); err != nil {
		return db, err
	}

	db = db.Scopes(
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
	)

	if filter.Not != nil {
		db = db.Scopes(func(d *gorm.DB) *gorm.DB {
			return d.Not(UUID(d, field, *filter.Not))
		})
	}

	if filter.And != nil {
		for _, and := range *filter.And {
			_filter := and

			db = db.Scopes(func(d *gorm.DB) *gorm.DB {
				return d.Where(UUID(d, field, _filter))
			})
		}
	}

	if filter.Or != nil {
		for _, or := range *filter.Or {
			_filter := or

			db = db.Scopes(func(d *gorm.DB) *gorm.DB {
				return d.Or(UUID(d, field, _filter))
			})
		}
	}

	return db, nil
}
