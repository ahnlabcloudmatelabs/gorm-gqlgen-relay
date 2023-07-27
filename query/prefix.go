package query

import (
	"fmt"
)

func HasPrefix(field string, value *string, queryString *string, values *[]any) {
	if value == nil {
		return
	}

	queryAppend(queryString, fmt.Sprintf("%s LIKE ?", field))
	*values = append(*values, fmt.Sprintf("%s%%", *value))
}

func HasSuffix(field string, value *string, queryString *string, values *[]any) {
	if value == nil {
		return
	}

	queryAppend(queryString, fmt.Sprintf("%s LIKE ?", field))
	*values = append(*values, fmt.Sprintf("%%%s", *value))
}
