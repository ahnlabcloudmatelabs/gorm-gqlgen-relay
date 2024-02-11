package where

import (
	"fmt"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/utils"
	"gorm.io/gorm"
)

func Traverse(db *gorm.DB, where Where) *gorm.DB {
	stmt := db.Where(where.Query, where.Args...)

	if where.Not != nil {
		stmt = stmt.Not(where.Not.Query, where.Not.Args...)
	}

	for _, and := range where.And {
		stmt = stmt.Where(Traverse(db, and))
	}

	for _, or := range where.Or {
		stmt = stmt.Or(Traverse(db, or))
	}

	return stmt
}

func traverse(dialector, table string, tables *map[string]string, schema *string, input map[string]any) (where Where) {
	for key, value := range input {
		if value == nil {
			continue
		}

		if key == "and" {
			for _, v := range value.([]any) {
				where.And = append(where.And, traverse(dialector, table, tables, schema, v.(map[string]any)))
			}
			continue
		}

		if key == "or" {
			for _, v := range value.([]any) {
				where.Or = append(where.Or, traverse(dialector, table, tables, schema, v.(map[string]any)))
			}
			continue
		}

		if key == "not" {
			where.Not = utils.ToPointer(traverse(dialector, table, tables, schema, value.(map[string]any)))
			continue
		}

		prefix := ""

		if table != "" {
			if schema != nil {
				prefix = "\"" + *schema + "\"."
			}

			prefix += "\"" + table + "\"" + "."
		}

		if tables != nil {
			for k, v := range *tables {
				if k == key {
					if schema != nil {
						prefix = "\"" + *schema + "\"."
					}

					prefix += "\"" + v + "\"" + "."
					break
				}
			}
		}
		key = "\"" + key + "\""
		_value, ok := value.(map[string]any)

		if !ok {
			fmt.Println("Conversion failed: value is not of type map[string]any")
			return
		}

		nonNilValues := map[string]any{}

		for k, v := range _value {
			if v == nil {
				continue
			}
			nonNilValues[k] = v
		}

		query, args := filter(dialector, prefix+key, nonNilValues)
		where.Query = utils.AppendQuery(where.Query, query)
		where.Args = append(where.Args, args...)
	}
	return
}
