package where

import "github.com/cloudmatelabs/gorm-gqlgen-relay/utils"

func HasKey(input any, key string) (bool, error) {
	filter, err := utils.ConvertToMapAny(input)
	if err != nil {
		return false, err
	}

	for k, v := range filter {
		if k == key {
			return true, nil
		}

		if k == "and" {
			for _, v := range v.([]any) {
				hasKey, err := HasKey(v, key)
				if err != nil {
					return false, err
				}

				if hasKey {
					return true, nil
				}
			}
		}

		if k == "or" {
			for _, v := range v.([]any) {
				hasKey, err := HasKey(v, key)
				if err != nil {
					return false, err
				}

				if hasKey {
					return true, nil
				}
			}
		}

		if k == "not" {
			hasKey, err := HasKey(v, key)
			if err != nil {
				return false, err
			}

			if hasKey {
				return true, nil
			}
		}
	}

	return false, nil
}
