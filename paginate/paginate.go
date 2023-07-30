package paginate

import (
	"github.com/cloudmatelabs/gorm-gqlgen-relay/cursor"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/edge"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/order"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/utils"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/where"
	"gorm.io/gorm"
)

type Option struct {
	First      *int
	Last       *int
	After      *string
	Before     *string
	Table      *string
	PrimaryKey string
}

type Connection[Model any] struct {
	TotalCount int64               `json:"totalCount"`
	Edges      []*edge.Edge[Model] `json:"edges"`
	PageInfo   *PageInfo           `json:"pageInfo"`
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor,omitempty"`
	EndCursor       *string `json:"endCursor,omitempty"`
}

func Paginate[Model any](db *gorm.DB, _where any, _orderBy any, option Option) (*Connection[Model], error) {
	if err := validation(option.First, option.Last, option.After, option.Before); err != nil {
		return nil, err
	}

	orderBy, err := utils.ConvertToMap(_orderBy)
	if err != nil {
		return nil, err
	}

	w, err := where.Do(_where)
	if err != nil {
		return nil, err
	}

	stmt := where.Traverse(db, w)
	var totalCount int64
	var model Model
	stmt.Model(&model).Count(&totalCount)

	stmt, err = setAfter(stmt, option.After, orderBy, option.PrimaryKey)
	if err != nil {
		return nil, err
	}
	stmt, err = setBefore(stmt, option.Before, orderBy, option.PrimaryKey)
	if err != nil {
		return nil, err
	}

	if queries, args, err := cursor.After(option.After, orderBy, option.PrimaryKey); err != nil {
		return nil, err
	} else {
		stmt = stmt.Where(queries, args...)
	}

	if queries, args, err := cursor.Before(option.Before, orderBy, option.PrimaryKey); err != nil {
		return nil, err
	} else {
		stmt = stmt.Where(queries, args...)
	}

	orders, err := order.By(_orderBy, option.Last != nil)
	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		stmt = stmt.Order(order)
	}

	_limit := limit(option.First, option.Last)
	if _limit != nil {
		stmt = stmt.Limit(*_limit)
	}

	var rows []*Model
	if err := stmt.Find(&rows).Error; err != nil {
		return nil, err
	}

	edges, err := edge.Convert(rows, utils.Keys(orderBy), option.PrimaryKey)
	if err != nil {
		return nil, err
	}

	pageInfo := PageInfo{
		HasNextPage: func() bool {
			if option.First == nil {
				return false
			}

			return totalCount > 0 && totalCount > int64(*option.First)
		}(),
		HasPreviousPage: func() bool {
			if option.Last == nil {
				return false
			}

			return totalCount > 0 && totalCount > int64(*option.Last)
		}(),
	}

	if len(edges) > 0 {
		pageInfo.StartCursor = &edges[0].Cursor
		pageInfo.EndCursor = &edges[len(edges)-1].Cursor
	}

	return &Connection[Model]{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo:   &pageInfo,
	}, nil
}

func setAfter(db *gorm.DB, after *string, orderBy map[string]any, primaryKey string) (*gorm.DB, error) {
	if after == nil {
		return db, nil
	}

	queries, args, err := cursor.After(after, orderBy, primaryKey)
	if err != nil {
		return db, err
	}

	return db.Where(queries, args...), nil
}

func setBefore(db *gorm.DB, before *string, orderBy map[string]any, primaryKey string) (*gorm.DB, error) {
	if before == nil {
		return db, nil
	}

	queries, args, err := cursor.Before(before, orderBy, primaryKey)
	if err != nil {
		return db, err
	}

	return db.Where(queries, args...), nil
}
