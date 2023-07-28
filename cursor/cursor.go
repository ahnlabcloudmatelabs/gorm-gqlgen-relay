package cursor

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/order"
)

func Set[Row, OrderBy, Edge any](rows []*Row, orderBy []*OrderBy, edges *[]*Edge) (*string, *string) {
	_orderBy := order.ParseOrderBy(orderBy)
	_edges := []map[string]interface{}{}

	for _, row := range rows {
		val := reflect.ValueOf(row).Elem()
		cursor := generateCursor(_orderBy, val)
		data, _ := json.Marshal(&cursor)
		_edges = append(_edges, map[string]interface{}{
			"node":   row,
			"cursor": base64.StdEncoding.EncodeToString(data),
		})
	}

	edgeData, _ := json.Marshal(&_edges)
	json.Unmarshal(edgeData, &edges)

	startCursor, endCursor := startEndCursor(_edges)
	return startCursor, endCursor
}

func generateCursor(orderBy []map[string]string, val reflect.Value) []interface{} {
	if len(orderBy) == 0 {
		return []interface{}{val.Type().Field(0).Name}
	}

	cursor := []interface{}{}

	for _, order := range orderBy {
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)

			if strings.ToLower(field.Name) == strings.ReplaceAll(order["field"], "_", "") {
				cursor = append(cursor, val.Field(i).Interface())
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
