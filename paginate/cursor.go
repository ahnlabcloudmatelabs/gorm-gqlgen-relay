package paginate

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/interfaces"
	"gorm.io/gorm"
)

func Cursor(db *gorm.DB, after *string, before *string, orderBy []map[string]string) (*gorm.DB, error) {
	if after == nil && before == nil {
		return db, nil
	}

	afterCursor, err := decodeCursor(after)
	if err != nil {
		return nil, err
	}

	beforeCursor, err := decodeCursor(before)
	if err != nil {
		return nil, err
	}

	if len(orderBy) == 0 {
		db.Scopes(func(d *gorm.DB) *gorm.DB {
			if after != nil {
				d = d.Where("id > ?", afterCursor[0])
			}
			if before != nil {
				d = d.Where("id < ?", beforeCursor[0])
			}

			return d
		})

		return db, nil
	}

	db = db.Scopes(func(d *gorm.DB) *gorm.DB {
		for i, order := range orderBy {
			if order["direction"] == "ASC" {
				if after != nil {
					d = d.Where(fmt.Sprintf("%s > ?", order["field"]), afterCursor[i])
				}
				if before != nil {
					d = d.Where(fmt.Sprintf("%s < ?", order["field"]), beforeCursor[i])
				}
			} else {
				if after != nil {
					d = d.Where(fmt.Sprintf("%s < ?", order["field"]), afterCursor[i])
				}
				if before != nil {
					d = d.Where(fmt.Sprintf("%s > ?", order["field"]), beforeCursor[i])
				}
			}
		}

		return d
	})

	return db, nil
}

func decodeCursor(data *string) ([]interface{}, error) {
	if data == nil {
		return nil, nil
	}

	cursorData, err := base64.StdEncoding.DecodeString(*data)
	if err != nil {
		return nil, err
	}
	cursor := []interface{}{}

	err = json.Unmarshal(cursorData, &cursor)
	return cursor, err
}

func StartCursor[Model any](edges []*interfaces.Edge[Model]) *string {
	if len(edges) == 0 {
		return nil
	}

	return &edges[0].Cursor
}

func EndCursor[Model any](edges []*interfaces.Edge[Model]) *string {
	if len(edges) == 0 {
		return nil
	}

	return &edges[len(edges)-1].Cursor
}
