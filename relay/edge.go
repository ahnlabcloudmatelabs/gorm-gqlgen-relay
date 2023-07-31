package relay

import "github.com/cloudmatelabs/gorm-gqlgen-relay/cursor"

func convertToEdge[T any](rows []*T, fields []string, primaryKey string) ([]*Edge[T], error) {
	edges := make([]*Edge[T], 0, len(rows))

	for _, row := range rows {
		cursor, err := cursor.Create(row, fields, primaryKey)
		if err != nil {
			return nil, err
		}

		edges = append(edges, &Edge[T]{
			Cursor: cursor,
			Node:   row,
		})
	}

	return edges, nil
}
