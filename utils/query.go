package utils

func AppendQuery(acc, query string) string {
	if acc == "" {
		return query
	}

	return acc + " AND " + query
}
