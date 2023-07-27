package query

func queryAppend(queryString *string, q string) {
	if *queryString != "" {
		*queryString += " AND "
	}

	*queryString += q
}
