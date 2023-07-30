package order

import "github.com/cloudmatelabs/gorm-gqlgen-relay/utils"

func traverse(filter map[string]any, reverse bool) (orders []string) {
	for field, direction := range filter {
		if reverse {
			direction = utils.ReverseDirection(direction.(string))
		}
		orders = append(orders, field+" "+direction.(string))
	}
	return
}
