package relay

import (
	"encoding/json"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/count"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/cursor"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/filters"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/interfaces"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/order"
	"gorm.io/gorm"
)

type ResolverProps struct {
	DB      *gorm.DB
	First   *int
	Last    *int
	After   *string
	Before  *string
	OrderBy any
	Filter  any
}

func Resolver[Model any](props ResolverProps, model *Model) (connection *interfaces.Connection[Model], err error) {
	if props.Filter != nil {
		filter := parseFilter(props.Filter)

		props.DB, _ = filters.Where(props.DB, &filter)

		if filter["and"] != nil {
			props.DB, _ = filters.And(props.DB, filter["and"])
		}
		if filter["or"] != nil {
			props.DB, _ = filters.Or(props.DB, filter["or"])
		}
		if filter["not"] != nil {
			props.DB, _ = filters.Not(props.DB, filter["not"])
		}
	}

	orderBy := parseOrderBy(props.OrderBy)
	props.DB, _ = cursor.After(props.DB, props.After, orderBy)
	props.DB, _ = cursor.Before(props.DB, props.Before, orderBy)
	totalCount := count.Count(props.DB, model)

	// props.DB = relay.SetFirst(props.DB, props.First)
	props.DB = order.OrderBy(props.DB, orderBy, "posts")

	var currentCount int64
	var rows []*Model
	err = props.DB.Model(model).Count(&currentCount).Find(&rows).Error
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

func parseFilter(input any) (filter map[string]any) {
	if input == nil {
		return nil
	}

	data, _ := json.Marshal(input)
	json.Unmarshal(data, &filter)
	return
}

func parseOrderBy(input any) (orderBy []map[string]string) {
	if input == nil {
		return nil
	}

	data, _ := json.Marshal(input)
	json.Unmarshal(data, &orderBy)
	return
}
