package query

func IsNull(field string, value *bool, queryString *string) {
	if value == nil {
		return
	}

	if *value {
		queryAppend(queryString, field+" IS NULL")
		return
	}

	queryAppend(queryString, field+" IS NOT NULL")
}

func IsNotNull(field string, value *bool, queryString *string) {
	if value == nil {
		return
	}

	if *value {
		queryAppend(queryString, field+" IS NOT NULL")
		return
	}

	queryAppend(queryString, field+" IS NULL")
}
