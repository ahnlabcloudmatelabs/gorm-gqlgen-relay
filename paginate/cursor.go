package paginate

import (
	"reflect"
	"strings"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/utils"
)

func createCursor[T any](row *T, fields []string, idColumn string) (string, error) {
	reflectRow := reflect.ValueOf(*row)

	if len(fields) == 0 {
		return utils.MapToBase64(map[string]any{
			idColumn: reflectRow.FieldByNameFunc(fieldWithColumnIsEqual(idColumn)).String(),
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
		lower := strings.ToLower(key)
		if lower == field {
			return true
		}

		return strings.ReplaceAll(lower, "_", "") == field
	}
}
