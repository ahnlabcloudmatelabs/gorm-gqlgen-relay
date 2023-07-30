package paginate

type Edge[T any] struct {
	Cursor string
	Node   T
}

func rowsToEdges[T any](rows []*T, fields []string, idColumn string) ([]Edge[T], error) {
	edges := make([]Edge[T], 0, len(rows))

	for _, row := range rows {
		cursor, err := createCursor(row, fields, idColumn)
		if err != nil {
			return nil, err
		}

		edges = append(edges, Edge[T]{
			Cursor: cursor,
			Node:   *row,
		})
	}

	return edges, nil
}
