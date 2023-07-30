package order

import "github.com/cloudmatelabs/gorm-gqlgen-relay/utils"

func traverse(filter map[string]any, reverse bool) (query string) {
	for field, direction := range filter {
		if reverse {
			direction = utils.ReverseDirection(direction.(string))
		}
		query = utils.AppendOrder(query, field, direction.(string))
	}
	return
}
