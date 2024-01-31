package cursor

import (
	"reflect"
	"strings"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/utils"
)

func Create[T any](row *T, fields []string, primaryKey string) (string, error) {
	reflectRow := reflect.ValueOf(*row)

	if len(fields) == 0 {
		return utils.MapToBase64(map[string]any{
			primaryKey: reflectRow.FieldByNameFunc(fieldWithColumnIsEqual(primaryKey)).Interface(),
		})
	}

	cursor := map[string]any{}

	for _, field := range fields {
		cursor[field] = reflectRow.FieldByNameFunc(fieldWithColumnIsEqual(field)).Interface()
	}

	return utils.MapToBase64(cursor)
}

func fieldWithColumnIsEqual(field string) func(key string) bool {
	return func(key string) bool {
		return strings.ToLower(key) == strings.ReplaceAll(strings.ToLower(field), "_", "")
	}
}
