package paginate

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/interfaces"
)

func Edges[Model any](rows []*Model, orderBy []map[string]string) ([]*interfaces.Edge[Model], error) {
	edges := []*interfaces.Edge[Model]{}

	for _, row := range rows {
		cursor := generateCursor(orderBy, reflect.ValueOf(row).Elem())
		data, err := json.Marshal(&cursor)
		if err != nil {
			return nil, err
		}

		edges = append(edges, &interfaces.Edge[Model]{
			Node:   row,
			Cursor: base64.StdEncoding.EncodeToString(data),
		})
	}

	return edges, nil
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
