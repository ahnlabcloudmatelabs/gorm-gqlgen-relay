package relay

import (
	"github.com/cloudmatelabs/gorm-gqlgen-relay/filters"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/interfaces"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/paginate"
	"gorm.io/gorm"
)

type ResolverProps struct {
	First   *int
	Last    *int
	After   *string
	Before  *string
	OrderBy any
	Filter  any
}

func Resolver[Model any](db *gorm.DB, model *Model, option ResolverProps) (*interfaces.Connection[Model], error) {
	paginateOptions := paginate.Options{
		First:  option.First,
		Last:   option.Last,
		After:  option.After,
		Before: option.Before,
	}

	db = filters.Do(db, option.Filter)
	connection, err := paginate.Paginate(db, *model, option.OrderBy, paginateOptions)

	return connection, err
}
