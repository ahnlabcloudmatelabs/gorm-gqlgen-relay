package order

import "github.com/cloudmatelabs/gorm-gqlgen-relay/utils"

func By(table string, input any, reverse bool) ([]string, error) {
	filter, err := utils.ConvertToMap(input)
	if err != nil {
		return nil, err
	}

	prefix := ""
	if table != "" {
		prefix = table + "."
	}

	return traverse(prefix, filter, reverse), nil
}
