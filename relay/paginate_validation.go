package relay

import (
	"encoding/base64"
	"fmt"
)

func validation(first, last *int, before, after *string) error {
	if first != nil && last != nil {
		return fmt.Errorf("passing both `first` and `last` to paginate a connection is not supported")
	}

	if first != nil && *first < 0 {
		return fmt.Errorf("`first` on a connection cannot be less than zero")
	}

	if last != nil && *last < 0 {
		return fmt.Errorf("`last` on a connection cannot be less than zero")
	}

	if after != nil {
		if _, err := base64.StdEncoding.DecodeString(*after); err != nil {
			return fmt.Errorf("invalid `after` cursor: `%s`", *after)
		}
	}

	if before != nil {
		if _, err := base64.StdEncoding.DecodeString(*before); err != nil {
			return fmt.Errorf("invalid `before` cursor: `%s`", *before)
		}
	}

	return nil
}
