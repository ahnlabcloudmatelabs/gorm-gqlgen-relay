package query

import (
	"fmt"
	"strings"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/utils"
	"gorm.io/gorm"
)

var (
	POSTGRES_FLATTEN_OPTIONS = utils.FlattenOptions{
		Delimiter: "->",
		ListStart: "->",
		ListEnd:   "->",
		KeyWrap:   "'",
	}
	SQL_FLATTEN_OPTIONS = utils.FlattenOptions{
		Delimiter: ".",
		ListStart: "[",
		ListEnd:   "]",
	}
)

func MapEqual(field string, input *map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	if input == nil {
		return self()
	}

	return func(db *gorm.DB) *gorm.DB {
		var scope func(db *gorm.DB) *gorm.DB

		switch db.Dialector.Name() {
		case "postgres":
			scope = gatherScopes("(%s->%s)::jsonb = '?'", field, input, "", POSTGRES_FLATTEN_OPTIONS)
		case "sqlserver":
			scope = gatherScopes("JSON_VALUE(%s, '%s') = '?'", field, input, "$.", SQL_FLATTEN_OPTIONS)
		default:
			scope = gatherScopes("JSON_EXTRACT(%s, '%s') = '?'", field, input, "$.", SQL_FLATTEN_OPTIONS)
		}

		return db.Scopes(scope)
	}
}

func MapNotEqual(field string, input *map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	if input == nil {
		return self()
	}

	return func(db *gorm.DB) *gorm.DB {
		var scope func(db *gorm.DB) *gorm.DB

		switch db.Dialector.Name() {
		case "postgres":
			scope = gatherScopes("(%s->%s)::jsonb != '?'", field, input, "", POSTGRES_FLATTEN_OPTIONS)
		case "sqlserver":
			scope = gatherScopes("JSON_VALUE(%s, '%s') != '?'", field, input, "$.", SQL_FLATTEN_OPTIONS)
		default:
			scope = gatherScopes("JSON_EXTRACT(%s, '%s') != '?'", field, input, "$.", SQL_FLATTEN_OPTIONS)
		}

		return db.Scopes(scope)
	}
}

func MapEqualFold(field string, input *map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	if input == nil {
		return self()
	}

	return func(db *gorm.DB) *gorm.DB {
		var scope func(db *gorm.DB) *gorm.DB

		switch db.Dialector.Name() {
		case "postgres":
			scope = gatherScopes("LOWER((%s->%s)::text) = '?'", field, input, "", POSTGRES_FLATTEN_OPTIONS)
		case "sqlserver":
			scope = gatherScopes("LOWER(JSON_VALUE(%s, '%s')) = '?'", field, input, "$.", SQL_FLATTEN_OPTIONS)
		default:
			scope = gatherScopes("LOWER(JSON_EXTRACT(%s, '%s')) = '?'", field, input, "$.", SQL_FLATTEN_OPTIONS)
		}

		return db.Scopes(scope)
	}
}

func gatherScopes(query string, field string, input *map[string]interface{}, prefix string, options utils.FlattenOptions) func(db *gorm.DB) *gorm.DB {
	var value map[string]interface{}
	utils.Flatten(prefix, *input, value, options)

	return func(db *gorm.DB) *gorm.DB {
		for k, v := range value {
			db = db.Where(fmt.Sprintf(query, field, k), strings.ToLower(v.(string)))
		}
		return db
	}
}
