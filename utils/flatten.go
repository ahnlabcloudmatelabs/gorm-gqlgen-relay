package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type FlattenOptions struct {
	Delimiter string
	ListStart string
	ListEnd   string
	KeyWrap   string
}

func Flatten(prefix string, src map[string]interface{}, dest map[string]interface{}, options FlattenOptions) error {
	jsonByte, err := json.Marshal(src)
	if err != nil {
		return err
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(jsonByte, &jsonData)
	if err != nil {
		return err
	}

	flatten(prefix, jsonData, dest, options)
	return nil
}

func flatten(prefix string, src map[string]interface{}, dest map[string]interface{}, options FlattenOptions) {
	for key, value := range src {
		switch reflect.TypeOf(value).Kind() {

		case reflect.Slice:
			for index := 0; index < len(value.([]interface{})); index++ {
				flatten(
					fmt.Sprintf("%s%d%s", prefix+options.KeyWrap+key+options.KeyWrap+options.ListStart, index, options.ListEnd+options.Delimiter),
					value.([]interface{})[index].(map[string]interface{}),
					dest,
					options,
				)
			}

		case reflect.Map:
			flatten(
				prefix+options.KeyWrap+key+options.KeyWrap+options.Delimiter,
				value.(map[string]interface{}),
				dest,
				options,
			)

		default:
			dest[prefix+key] = value
		}
	}
}
