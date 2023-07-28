package filters

import (
	"reflect"

	"github.com/cloudmatelabs/gorm-gqlgen-relay/query"
)

func createQuery(field string, filter Filter[any]) (queryString string, values []any) {
	query.Equal(field, filter.Equal, &queryString, &values)
	query.NotEqual(field, filter.NotEqual, &queryString, &values)
	query.In(field, filter.In, &queryString, &values)
	query.NotIn(field, filter.NotIn, &queryString, &values)
	query.GreaterThan(field, filter.GreaterThan, &queryString, &values)
	query.GreaterThanOrEqual(field, filter.GreaterThanOrEqual, &queryString, &values)
	query.LessThan(field, filter.LessThan, &queryString, &values)
	query.LessThanOrEqual(field, filter.LessThanOrEqual, &queryString, &values)
	query.IsNull(field, filter.IsNull, &queryString)
	query.IsNotNull(field, filter.IsNotNull, &queryString)

	appendStringIDQuery(field, filter.EqualFold, &queryString, &values, query.EqualFold)
	appendStringIDQuery(field, filter.Contains, &queryString, &values, query.Contains)
	appendStringIDQuery(field, filter.ContainsFold, &queryString, &values, query.ContainsFold)
	appendStringIDQuery(field, filter.HasPrefix, &queryString, &values, query.HasPrefix)
	appendStringIDQuery(field, filter.HasSuffix, &queryString, &values, query.HasSuffix)
	return
}

func appendStringIDQuery(field string, input *any, queryString *string, values *[]any, callback func(string, *string, *string, *[]any)) {
	if input == nil {
		return
	}

	if reflect.ValueOf(*input).Kind() != reflect.String {
		return
	}

	value := (*input).(string)
	callback(field, &value, queryString, values)
}
