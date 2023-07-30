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

func ReverseInequality(inequality string) string {
	if inequality == ">" {
		return "<"
	}

	if inequality == "<" {
		return ">"
	}

	if inequality == ">=" {
		return "<="
	}

	if inequality == "<=" {
		return ">="
	}

	return inequality
}
