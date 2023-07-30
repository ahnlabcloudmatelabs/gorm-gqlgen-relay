package order

import "github.com/cloudmatelabs/gorm-gqlgen-relay/utils"

func By(input any, reverse bool) ([]string, error) {
	filter, err := utils.ConvertToMap(input)
	if err != nil {
		return nil, err
	}

	return traverse(filter, reverse), nil
}
