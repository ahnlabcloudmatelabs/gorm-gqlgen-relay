package cursor

import (
	"fmt"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/utils"
)

func After(after *string, orderBy map[string]any, primaryKey string) (queries []string, args []any, err error) {
	return loadCursor(">")(after, orderBy, primaryKey)
}

func Before(before *string, orderBy map[string]any, primaryKey string) (queries []string, args []any, err error) {
	return loadCursor("<")(before, orderBy, primaryKey)
}

func loadCursor(defaultInequality string) func(cursor *string, orderBy map[string]any, primaryKey string) (queries []string, args []any, err error) {
	return func(cursor *string, orderBy map[string]any, primaryKey string) (queries []string, args []any, err error) {
		if cursor == nil {
			return
		}

		_cursor, err := utils.ParseCursor(*cursor)
		if err != nil {
			return
		}

		if orderBy == nil {
			orderBy = map[string]any{primaryKey: "ASC"}
		}

		if !utils.SameKeys(_cursor, orderBy) {
			return []string{}, []any{}, fmt.Errorf("input error: incorrect cursor given, deserialised base64 cursor map keys must match orderBy map keys, cursor keys: %v vs orderBy keys: %v", utils.GetMapKeys(_cursor), utils.GetMapKeys(orderBy))
		}

		for field, value := range _cursor {
			if v, ok := orderBy[field]; ok {
				direction := v.(string)
				queries = append(queries, fmt.Sprintf("\"%s\" %s ?", field, inequality(defaultInequality, direction)))
				args = append(args, value)
			}
		}
		return
	}
}

func inequality(input, direction string) string {
	if direction == "DESC" || direction == "desc" {
		return utils.ReverseInequality(input)
	}

	return input
}
