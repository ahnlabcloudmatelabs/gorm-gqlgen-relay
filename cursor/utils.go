package cursor

import (
	"encoding/base64"
	"encoding/json"
)

func decodeCursor(base64Cursor *string) ([]interface{}, error) {
	cursorData, err := base64.StdEncoding.DecodeString(*base64Cursor)
	if err != nil {
		return nil, err
	}
	cursor := []interface{}{}

	err = json.Unmarshal(cursorData, &cursor)
	return cursor, err
}
