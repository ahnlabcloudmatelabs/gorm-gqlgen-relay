package query

func Equal[T any](field string, value *T, queryString *string, values *[]any) {
	if value == nil {
		return
	}

	queryAppend(queryString, field+" = ?")
	*values = append(*values, *value)
}

func NotEqual[T any](field string, value *T, queryString *string, values *[]any) {
	if value == nil {
		return
	}

	queryAppend(queryString, field+" <> ?")
	*values = append(*values, *value)
}

func EqualFold(field string, value *string, queryString *string, values *[]any) {
	if value == nil {
		return
	}

	queryAppend(queryString, "LOWER("+field+") = LOWER(?)")
	*values = append(*values, *value)
}
