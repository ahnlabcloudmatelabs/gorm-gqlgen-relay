package relay

func hasNextPage[T any](totalCount int64, first *int, edges []*Edge[T]) bool {
	currentCount := len(edges)

	if first == nil {
		return totalCount > int64(currentCount)
	}

	return currentCount > *first
}

func hasPreviousPage[T any](totalCount int64, last *int, edges []*Edge[T]) bool {
	currentCount := len(edges)

	if last == nil {
		return totalCount > int64(currentCount)
	}

	return currentCount > *last
}
