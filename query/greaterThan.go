package query

func GreaterThan[T any](field string, value *T, queryString *string, values *[]any) {
	if value == nil {
		return
	}

	queryAppend(queryString, field+" > ?")
	*values = append(*values, *value)
}

func GreaterThanOrEqual[T any](field string, value *T, queryString *string, values *[]any) {
	if value == nil {
		return
	}

	queryAppend(queryString, field+" >= ?")
	*values = append(*values, *value)
}
