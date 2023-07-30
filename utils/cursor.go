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
