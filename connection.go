package relay

import (
	"github.com/cloudmatelabs/gorm-gqlgen-relay/count"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/cursor"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/filters"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/interfaces"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/order"
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

func Resolver[Model any](db *gorm.DB, model *Model, option ResolverProps) (connection *interfaces.Connection[Model], err error) {
	db = filters.Do(db, option.Filter)
	orderBy := order.ParseOrderBy(option.OrderBy)
	db, _ = cursor.After(db, option.After, orderBy)
	db, _ = cursor.Before(db, option.Before, orderBy)
	totalCount := count.Count(db, model)

	// db = relay.SetFirst(db, props.First)
	db = order.OrderBy(db, orderBy, "posts")

	var currentCount int64
	var rows []*Model
	err = db.Model(model).Count(&currentCount).Find(&rows).Error
	if err != nil {
		return
	}

	startCursor, endCursor, edges := cursor.Set(rows, orderBy)

	connection = &interfaces.Connection[Model]{
		TotalCount: int(totalCount),
		Edges:      edges,
		PageInfo: &interfaces.PageInfo{
			HasPreviousPage: currentCount > 0 && currentCount < totalCount,
			HasNextPage:     currentCount < totalCount && endCursor != nil,
			StartCursor:     startCursor,
			EndCursor:       endCursor,
		},
	}

	return
}
