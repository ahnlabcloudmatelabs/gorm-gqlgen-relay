package paginate

import (
	"github.com/cloudmatelabs/gorm-gqlgen-relay/interfaces"
	"gorm.io/gorm"
)

func Paginate[Model any](db *gorm.DB, model Model, orderBy any, options Options) (*interfaces.Connection[Model], error) {
	if err := Validation(options); err != nil {
		return nil, err
	}

	db = db.Model(model)

	totalCount, err := TotalCount(db, model)
	if err != nil {
		return nil, err
	}

	_orderBy, err := ParseOrderBy(orderBy)
	if err != nil {
		return nil, err
	}

	db = OrderBy(db, _orderBy, options.Table, options.Last != nil)
	db, err = Cursor(db, options.After, options.Before, _orderBy)
	if err != nil {
		return nil, err
	}

	db = Limit(db, options.First, options.Last)

	var rows []*Model
	if err := db.Find(&rows).Error; err != nil {
		return nil, err
	}

	edges, err := Edges(rows, _orderBy)
	if err != nil {
		return nil, err
	}

	return &interfaces.Connection[Model]{
		TotalCount: int(totalCount),
		Edges:      edges,
		PageInfo: &interfaces.PageInfo{
			StartCursor: StartCursor(edges),
			EndCursor:   EndCursor(edges),
			HasPreviousPage: func() bool {
				if options.Last == nil {
					return false
				}
				return totalCount > 0 && totalCount > int64(*options.Last)
			}(),
			HasNextPage: func() bool {
				if options.First == nil {
					return false
				}
				return totalCount > 0 && totalCount > int64(*options.First)
			}(),
		},
	}, nil
}
