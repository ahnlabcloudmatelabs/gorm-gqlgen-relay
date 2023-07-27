package query

func LessThan[T any](field string, value *T, queryString *string, values *[]any) {
	if value == nil {
		return
	}

	queryAppend(queryString, field+" < ?")
	*values = append(*values, *value)
}

func LessThanOrEqual[T any](field string, value *T, queryString *string, values *[]any) {
	if value == nil {
		return
	}

	queryAppend(queryString, field+" <= ?")
	*values = append(*values, *value)
}
