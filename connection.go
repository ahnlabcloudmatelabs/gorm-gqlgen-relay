package relay

import (
	"encoding/json"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/count"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/cursor"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/filters"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/order"
	"gorm.io/gorm"
)

type Props struct {
	DB      *gorm.DB
	First   *int
	Last    *int
	After   *string
	Before  *string
	OrderBy any
	Filter  any
}

func Connection[Model, Connection any](props Props, model *Model, connection *Connection) {
	if props.Filter != nil {
		filter := parseFilter(props.Filter)

		props.DB, _ = filters.Where(props.DB, &filter)

		if filter["and"] != nil {
			props.DB, _ = filters.And(props.DB, filter["and"])
		}
		if filter["or"] != nil {
			props.DB, _ = filters.Or(props.DB, filter["or"])
		}
	}

	orderBy := parseOrderBy(props.OrderBy)
	props.DB, _ = cursor.After(props.DB, props.After, orderBy)
	props.DB, _ = cursor.Before(props.DB, props.Before, orderBy)

	totalCount := count.Count(props.DB, model)

	// props.DB = relay.SetFirst(props.DB, props.First)
	props.DB = order.OrderBy(props.DB, orderBy, "posts")

	var currentCount int64
	var rows []Model
	props.DB.Model(model).Count(&currentCount).Debug().Find(&rows)

	startCursor, endCursor, edges := cursor.Set(rows, orderBy)

	result := map[string]any{
		"totalCount": totalCount,
		"edges":      edges,
		"pageInfo": map[string]any{
			"hasPreviousPage": currentCount > 0 && currentCount < totalCount,
			"hasNextPage":     currentCount < totalCount && endCursor != nil,
			"startCursor":     startCursor,
			"endCursor":       endCursor,
		},
	}

	data, _ := json.Marshal(result)
	json.Unmarshal(data, connection)
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
