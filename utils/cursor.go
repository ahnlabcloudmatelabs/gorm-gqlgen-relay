package utils

import (
	"encoding/base64"
	"encoding/json"
)

func ParseCursor(src string) (map[string]any, error) {
	data, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return nil, err
	}

	cursor := map[string]any{}
	if err = json.Unmarshal(data, &cursor); err != nil {
		return nil, err
	}

	return cursor, nil
}

func SameKeys(map1, map2 map[string]interface{}) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key := range map1 {
		if _, ok := map2[key]; !ok {
			return false
		}
	}

	return true
}

func GetMapKeys(inputMap map[string]interface{}) []string {
	keys := make([]string, 0, len(inputMap))
	for key := range inputMap {
		keys = append(keys, key)
	}
	return keys
}
