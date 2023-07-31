package order

import "github.com/cloudmatelabs/gorm-gqlgen-relay/utils"

func traverse(prefix string, filter map[string]any, reverse bool) (orders []string) {
	for field, direction := range filter {
		if reverse {
			direction = utils.ReverseDirection(direction.(string))
		}
		orders = append(orders, prefix+field+" "+direction.(string))
	}
	return
}
