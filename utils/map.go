package utils

import "encoding/json"

func ConvertToMap(src any) (map[string]any, error) {
	data, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}

	dest := make(map[string]any)
	err = json.Unmarshal(data, &dest)
	return dest, err
}
