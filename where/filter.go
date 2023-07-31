package where

import (
	"fmt"
	"strings"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/utils"
)

func filter(dialector, column string, input map[string]any) (query string, args []any) {
	for key, value := range input {
		if strings.Contains((fmt.Sprintf("%T", value)), "map[string]") {
			return mapFilter(dialector, column, key, value.(map[string]any))
		}

		switch key {
		case "equal":
			query = utils.AppendQuery(query, column+" = ?")
			args = append(args, value)
		case "notEqual":
			query = utils.AppendQuery(query, column+" != ?")
			args = append(args, value)
		case "equalFold":
			query = utils.AppendQuery(query, "LOWER("+column+") = LOWER(?)")
			args = append(args, value)
		case "in":
			query = utils.AppendQuery(query, column+" IN (?)")
			args = append(args, value)
		case "notIn":
			query = utils.AppendQuery(query, column+" NOT IN (?)")
			args = append(args, value)
		case "contains":
			query = utils.AppendQuery(query, column+" LIKE ?")
			args = append(args, "%"+value.(string)+"%")
		case "containsFold":
			query = utils.AppendQuery(query, "LOWER("+column+") LIKE LOWER(?)")
			args = append(args, "%"+value.(string)+"%")
		case "gt":
			query = utils.AppendQuery(query, column+" > ?")
			args = append(args, value)
		case "gte":
			query = utils.AppendQuery(query, column+" >= ?")
			args = append(args, value)
		case "lt":
			query = utils.AppendQuery(query, column+" < ?")
			args = append(args, value)
		case "lte":
			query = utils.AppendQuery(query, column+" <= ?")
			args = append(args, value)
		case "hasPrefix":
			query = utils.AppendQuery(query, column+" LIKE ?")
			args = append(args, value.(string)+"%")
		case "hasSuffix":
			query = utils.AppendQuery(query, column+" LIKE ?")
			args = append(args, "%"+value.(string))
		case "isNull":
			if value.(bool) {
				query = utils.AppendQuery(query, column+" IS NULL")
			} else {
				query = utils.AppendQuery(query, column+" IS NOT NULL")
			}
		}
	}
	return
}
