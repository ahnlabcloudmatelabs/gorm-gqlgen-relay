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

	w, err := where.Do(_where)
	if err != nil {
		return nil, err
	}

	stmt := where.Traverse(db, w)
	var totalCount int64
	var model Model
	stmt.Model(&model).Count(&totalCount)

	orderBy, err := utils.ConvertToMap(_orderBy)
	if err != nil {
		return nil, err
	}

	stmt, err = setAfter(stmt, option.After, orderBy, option.PrimaryKey)
	if err != nil {
		return nil, err
	}
	stmt, err = setBefore(stmt, option.Before, orderBy, option.PrimaryKey)
	if err != nil {
		return nil, err
	}

	orders, err := order.By(_orderBy, option.Last != nil)
	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		stmt = stmt.Order(order)
	}

	stmt = limit(stmt, option.First, option.Last)

	var rows []*Model
	if err := stmt.Find(&rows).Error; err != nil {
		return nil, err
	}

	edges, err := edge.Convert(rows, utils.Keys(orderBy), option.PrimaryKey)
	if err != nil {
		return nil, err
	}

	pageInfo := &PageInfo{
		HasNextPage:     hasNextPage(totalCount, option.First, edges),
		HasPreviousPage: hasPreviousPage(totalCount, option.Last, edges),
	}

	if len(edges) > 0 {
		pageInfo.StartCursor = &edges[0].Cursor
		pageInfo.EndCursor = &edges[len(edges)-1].Cursor
	}

	return &Connection[Model]{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo:   pageInfo,
	}, nil
}

func setAfter(db *gorm.DB, after *string, orderBy map[string]any, primaryKey string) (*gorm.DB, error) {
	return setCursor(db, after, "after", orderBy, primaryKey)
}

func setBefore(db *gorm.DB, before *string, orderBy map[string]any, primaryKey string) (*gorm.DB, error) {
	return setCursor(db, before, "before", orderBy, primaryKey)
}

func setCursor(db *gorm.DB, cur *string, direction string, orderBy map[string]any, primaryKey string) (*gorm.DB, error) {
	if cur == nil {
		return db, nil
	}

	var queries []string
	var args []any
	var err error

	if direction == "after" {
		queries, args, err = cursor.After(cur, orderBy, primaryKey)
	} else {
		queries, args, err = cursor.Before(cur, orderBy, primaryKey)
	}

	if err != nil {
		return db, err
	}

	for i := range queries {
		db = db.Where(queries[i], args[i])
	}

	return db, err
}

func hasNextPage[T any](totalCount int64, first *int, edges []*edge.Edge[T]) bool {
	currentCount := len(edges)

	if first == nil {
		return totalCount > int64(currentCount)
	}

	return currentCount > *first
}

func hasPreviousPage[T any](totalCount int64, last *int, edges []*edge.Edge[T]) bool {
	currentCount := len(edges)

	if last == nil {
		return totalCount > int64(currentCount)
	}

	return currentCount > *last
}
