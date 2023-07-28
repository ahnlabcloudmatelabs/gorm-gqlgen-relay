package cursor

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"strings"
)

func Set[T any](rows []T, orderBy []map[string]string) (*string, *string, []map[string]any) {
	_edges := []map[string]interface{}{}

	for _, row := range rows {
		cursor := generateCursor(orderBy, reflect.ValueOf(row))
		data, _ := json.Marshal(&cursor)
		_edges = append(_edges, map[string]interface{}{
			"node":   row,
			"cursor": base64.StdEncoding.EncodeToString(data),
		})
	}

	edges := []map[string]any{}
	edgeData, _ := json.Marshal(&_edges)
	json.Unmarshal(edgeData, &edges)

	startCursor, endCursor := startEndCursor(_edges)
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

func startEndCursor(edges []map[string]interface{}) (startCursor, endCursor *string) {
	if len(edges) > 0 {
		_startCursor := edges[0]["cursor"].(string)
		_endCursor := edges[len(edges)-1]["cursor"].(string)

		startCursor = &_startCursor
		endCursor = &_endCursor
	}

	return
}
