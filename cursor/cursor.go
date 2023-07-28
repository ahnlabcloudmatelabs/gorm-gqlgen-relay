package cursor

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/interfaces"
)

func Set[T any](rows []*T, orderBy []map[string]string) (*string, *string, []*interfaces.Edge[T]) {
	edges := []*interfaces.Edge[T]{}

	for _, row := range rows {
		cursor := generateCursor(orderBy, reflect.ValueOf(row))
		data, _ := json.Marshal(&cursor)
		edges = append(edges, &interfaces.Edge[T]{
			Node:   row,
			Cursor: base64.StdEncoding.EncodeToString(data),
		})
	}

	startCursor, endCursor := startEndCursor(edges)
	return startCursor, endCursor, edges
}

func generateCursor(orderBy []map[string]string, row reflect.Value) []any {
	if len(orderBy) == 0 {
		return []any{row.Type().Field(0).Name}
	}

	cursor := []any{}

	for _, order := range orderBy {
		for i := 0; i < row.NumField(); i++ {
			field := row.Type().Field(i)

			if strings.ToLower(field.Name) == strings.ReplaceAll(order["field"], "_", "") || field.Name == order["field"] {
				cursor = append(cursor, row.Field(i).Interface())
				break
			}
		}
	}

	return cursor
}

func startEndCursor[T any](edges []*interfaces.Edge[T]) (startCursor, endCursor *string) {
	if len(edges) > 0 {
		_startCursor := edges[0].Cursor
		_endCursor := edges[len(edges)-1].Cursor

		startCursor = &_startCursor
		endCursor = &_endCursor
	}

	return
}
