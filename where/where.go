package where

import "github.com/cloudmatelabs/gorm-gqlgen-relay/utils"

type Where struct {
	And []Where
	Or  []Where
	Not *Where

	Query string
	Args  []any
}

func Do(input any) (Where, error) {
	filter, err := utils.ConvertToMap(input)
	if err != nil {
		return Where{}, err
	}

	return traverse(filter), nil
}
