package paginate

func limit(first, last *int) *int {
	if first != nil {
		return first
	}

	if last != nil {
		return last
	}

	return nil
}
