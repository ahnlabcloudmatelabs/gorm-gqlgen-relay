package order

import "github.com/cloudmatelabs/gorm-gqlgen-relay/utils"

func traverse(table string, tables *map[string]string, filter map[string]any, reverse bool) (orders []string) {
	for field, direction := range filter {
		if reverse {
			direction = utils.ReverseDirection(direction.(string))
		}

		prefix := ""

		if tables != nil {
			for k, v := range *tables {
				if k == field {
					prefix = v + "."
					break
				}
			}
		} else {
			if table != "" {
				prefix = table + "."
			}
		}

		orders = append(orders, prefix+field+" "+direction.(string))
	}
	return
}
