package where

import (
	"fmt"
	"strings"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/utils"
)

func mapFilter(dialector, column, key string, value map[string]any) (query string, args []any) {
	switch dialector {
	case "sqlserver":
		return mapFilterSqlServer(column, key, value)
	case "postgres":
		return mapFilterPostgres(column, key, value)
	default:
		return mapFilterMySQL(column, key, value)
	}
}

func mapFilterSqlServer(column, key string, value map[string]any) (query string, args []any) {
	field := value["key"].(string)

	switch key {
	case "equal":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_VALUE(%s, '$.\"%s\"') = ?", column, field))
		args = append(args, value["value"])
	case "notEqual":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_VALUE(%s, '$.\"%s\"') != ?", column, field))
		args = append(args, value["value"])
	case "equalFold":
		query = utils.AppendQuery(query, fmt.Sprintf("LOWER(JSON_VALUE(%s, '$.\"%s\"')) = LOWER(?)", column, field))
		args = append(args, value["value"])
	case "in":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_VALUE(%s, '$.\"%s\"') IN (?)", column, field))
		args = append(args, value["value"])
	case "notIn":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_VALUE(%s, '$.\"%s\"') NOT IN (?)", column, field))
		args = append(args, value["value"])
	case "contains":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_VALUE(%s, '$.\"%s\"') LIKE ?", column, field))
		args = append(args, "%"+value["value"].(string)+"%")
	case "containsFold":
		query = utils.AppendQuery(query, fmt.Sprintf("LOWER(JSON_VALUE(%s, '$.\"%s\"')) LIKE LOWER(?)", column, field))
		args = append(args, "%"+value["value"].(string)+"%")
	case "gt":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_VALUE(%s, '$.\"%s\"') > ?", column, field))
		args = append(args, value["value"])
	case "gte":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_VALUE(%s, '$.\"%s\"') >= ?", column, field))
		args = append(args, value["value"])
	case "lt":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_VALUE(%s, '$.\"%s\"') < ?", column, field))
		args = append(args, value["value"])
	case "lte":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_VALUE(%s, '$.\"%s\"') <= ?", column, field))
		args = append(args, value["value"])
	case "hasPrefix":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_VALUE(%s, '$.\"%s\"') LIKE ?", column, field))
		args = append(args, value["value"].(string)+"%")
	case "hasSuffix":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_VALUE(%s, '$.\"%s\"') LIKE ?", column, field))
		args = append(args, "%"+value["value"].(string))
	case "isNull":
		if value["value"].(bool) {
			query = utils.AppendQuery(query, fmt.Sprintf("JSON_VALUE(%s, '$.\"%s\"') IS NULL", column, field))
		} else {
			query = utils.AppendQuery(query, fmt.Sprintf("JSON_VALUE(%s, '$.\"%s\"') IS NOT NULL", column, field))
		}
	}
	return
}

func mapFilterPostgres(column, key string, value map[string]any) (query string, args []any) {
	field := ""
	split := strings.Split(value["key"].(string), ".")
	for i, v := range split {
		arrow := "->"
		if i == len(split)-1 {
			arrow = "->>"
		}

		field += fmt.Sprintf("%s'%s'", arrow, v)
	}

	switch key {
	case "equal":
		query = utils.AppendQuery(query, fmt.Sprintf("%s%s = ?", column, field))
		args = append(args, value["value"])
	case "notEqual":
		query = utils.AppendQuery(query, fmt.Sprintf("%s%s != ?", column, field))
		args = append(args, value["value"])
	case "equalFold":
		query = utils.AppendQuery(query, fmt.Sprintf("LOWER(%s%s) = LOWER(?)", column, field))
		args = append(args, value["value"])
	case "in":
		query = utils.AppendQuery(query, fmt.Sprintf("%s%s IN (?)", column, field))
		args = append(args, value["value"])
	case "notIn":
		query = utils.AppendQuery(query, fmt.Sprintf("%s%s NOT IN (?)", column, field))
		args = append(args, value["value"])
	case "contains":
		query = utils.AppendQuery(query, fmt.Sprintf("%s%s LIKE ?", column, field))
		args = append(args, "%"+value["value"].(string)+"%")
	case "containsFold":
		query = utils.AppendQuery(query, fmt.Sprintf("LOWER(%s%s) LIKE LOWER(?)", column, field))
		args = append(args, "%"+value["value"].(string)+"%")
	case "gt":
		query = utils.AppendQuery(query, fmt.Sprintf("%s%s > ?", column, field))
		args = append(args, value["value"])
	case "gte":
		query = utils.AppendQuery(query, fmt.Sprintf("%s%s >= ?", column, field))
		args = append(args, value["value"])
	case "lt":
		query = utils.AppendQuery(query, fmt.Sprintf("%s%s < ?", column, field))
		args = append(args, value["value"])
	case "lte":
		query = utils.AppendQuery(query, fmt.Sprintf("%s%s <= ?", column, field))
		args = append(args, value["value"])
	case "hasPrefix":
		query = utils.AppendQuery(query, fmt.Sprintf("%s%s LIKE ?", column, field))
		args = append(args, value["value"].(string)+"%")
	case "hasSuffix":
		query = utils.AppendQuery(query, fmt.Sprintf("%s%s LIKE ?", column, field))
		args = append(args, "%"+value["value"].(string))
	case "isNull":
		if value["value"].(bool) {
			query = utils.AppendQuery(query, fmt.Sprintf("%s%s IS NULL", column, field))
		} else {
			query = utils.AppendQuery(query, fmt.Sprintf("%s%s IS NOT NULL", column, field))
		}
	}
	return
}

func mapFilterMySQL(column, key string, value map[string]any) (query string, args []any) {
	field := value["key"].(string)

	switch key {
	case "equal":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_EXTRACT(%s, '$.\"%s\"') = ?", column, field))
		args = append(args, value["value"])
	case "notEqual":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_EXTRACT(%s, '$.\"%s\"') != ?", column, field))
		args = append(args, value["value"])
	case "equalFold":
		query = utils.AppendQuery(query, fmt.Sprintf("LOWER(JSON_EXTRACT(%s, '$.\"%s\"')) = LOWER(?)", column, field))
		args = append(args, value["value"])
	case "in":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_EXTRACT(%s, '$.\"%s\"') IN (?)", column, field))
		args = append(args, value["value"])
	case "notIn":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_EXTRACT(%s, '$.\"%s\"') NOT IN (?)", column, field))
		args = append(args, value["value"])
	case "contains":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_EXTRACT(%s, '$.\"%s\"') LIKE ?", column, field))
		args = append(args, "%"+value["value"].(string)+"%")
	case "containsFold":
		query = utils.AppendQuery(query, fmt.Sprintf("LOWER(JSON_EXTRACT(%s, '$.\"%s\"')) LIKE LOWER(?)", column, field))
		args = append(args, "%"+value["value"].(string)+"%")
	case "gt":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_EXTRACT(%s, '$.\"%s\"') > ?", column, field))
		args = append(args, value["value"])
	case "gte":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_EXTRACT(%s, '$.\"%s\"') >= ?", column, field))
		args = append(args, value["value"])
	case "lt":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_EXTRACT(%s, '$.\"%s\"') < ?", column, field))
		args = append(args, value["value"])
	case "lte":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_EXTRACT(%s, '$.\"%s\"') <= ?", column, field))
		args = append(args, value["value"])
	case "hasPrefix":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_EXTRACT(%s, '$.\"%s\"') LIKE ?", column, field))
		args = append(args, value["value"].(string)+"%")
	case "hasSuffix":
		query = utils.AppendQuery(query, fmt.Sprintf("JSON_EXTRACT(%s, '$.\"%s\"') LIKE ?", column, field))
		args = append(args, "%"+value["value"].(string))
	case "isNull":
		if value["value"].(bool) {
			query = utils.AppendQuery(query, fmt.Sprintf("JSON_EXTRACT(%s, '$.\"%s\"') IS NULL", column, field))
		} else {
			query = utils.AppendQuery(query, fmt.Sprintf("JSON_EXTRACT(%s, '$.\"%s\"') IS NOT NULL", column, field))
		}
	}
	return
}
