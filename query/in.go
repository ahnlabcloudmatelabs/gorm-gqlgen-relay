package query

func In[T any](field string, value *T, queryString *string, values *[]any) {
	if value == nil {
		return
	}

	queryAppend(queryString, field+" IN ?")
	*values = append(*values, *value)
}

func NotIn[T any](field string, value *T, queryString *string, values *[]any) {
	if value == nil {
		return
	}

	queryAppend(queryString, field+" NOT IN ?")
	*values = append(*values, *value)
}
