package where

import (
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

func traverse(input map[string]any) (where Where) {
	for key, value := range input {
		if key == "and" {
			for _, v := range value.([]any) {
				where.And = append(where.And, traverse(v.(map[string]any)))
			}
			continue
		}

		if key == "or" {
			for _, v := range value.([]any) {
				where.Or = append(where.Or, traverse(v.(map[string]any)))
			}
			continue
		}

		if key == "not" {
			where.Not = utils.ToPointer(traverse(value.(map[string]any)))
			continue
		}

		query, args := filter(key, value.(map[string]any))
		where.Query = utils.AppendQuery(where.Query, query)
		where.Args = append(where.Args, args...)
	}
	return
}
