package where

import "github.com/cloudmatelabs/gorm-gqlgen-relay/utils"

type Where struct {
	And []Where
	Or  []Where
	Not *Where

	Query string
	Args  []any
}

func Do(dialector, table string, tables *map[string]string, schema *string, input any) (Where, error) {
	filter, err := utils.ConvertToMap(input)
	if err != nil {
		return Where{}, err
	}

	return traverse(dialector, table, tables, schema, filter), nil
}
