package query

import (
	"fmt"
	"strings"
)

func Contains(field string, value *string, queryString *string, values *[]any) {
	if value == nil {
		return
	}

	queryAppend(queryString, fmt.Sprintf("%s LIKE ?", field))
	*values = append(*values, fmt.Sprintf("%%%s%%", *value))
}

func ContainsFold(field string, value *string, queryString *string, values *[]any) {
	if value == nil {
		return
	}

	queryAppend(queryString, fmt.Sprintf("LOWER(%s) LIKE ?", field))
	*values = append(*values, fmt.Sprintf("%%%s%%", strings.ToLower(*value)))
}
