package utils

import (
	"encoding/base64"
	"encoding/json"
)

func ConvertToMap(src any) (map[string]any, error) {
	data, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}

	dest := make(map[string]any)
	err = json.Unmarshal(data, &dest)
	return dest, err
}

func MapToBase64(src map[string]any) (string, error) {
	data, err := json.Marshal(src)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

func Keys(m map[string]any) []string {
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	return keys
}
