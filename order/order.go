package order

import "github.com/cloudmatelabs/gorm-gqlgen-relay/utils"

func By(table string, tables *map[string]string, input any, reverse bool) ([]string, error) {
	filter, err := utils.ConvertToMap(input)
	if err != nil {
		return nil, err
	}

	return traverse(table, tables, filter, reverse), nil
}

func HasKey(input any, key string) (bool, error) {
	filter, err := utils.ConvertToMapString(input)
	if err != nil {
		return false, err
	}

	for k := range filter {
		if k == key {
			return true, nil
		}
	}

	return false, nil
}
