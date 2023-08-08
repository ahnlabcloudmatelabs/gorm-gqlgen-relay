package order

import "github.com/cloudmatelabs/gorm-gqlgen-relay/utils"

func By(table string, tables *map[string]string, input any, reverse bool) ([]string, error) {
	filter, err := utils.ConvertToMap(input)
	if err != nil {
		return nil, err
	}

	return traverse(table, tables, filter, reverse), nil
}
