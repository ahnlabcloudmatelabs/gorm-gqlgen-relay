package utils

func AppendOrder(acc, field, direction string) string {
	if acc == "" {
		return field + " " + direction
	}

	return acc + ", " + field + " " + direction
}

func ReverseDirection(direction string) string {
	if direction == "ASC" || direction == "asc" {
		return "DESC"
	}

	return "ASC"
}
