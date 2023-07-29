package paginate

import (
	"encoding/base64"
	"fmt"
)

func Validation(options Options) error {
	if options.First != nil && options.Last != nil {
		return fmt.Errorf("passing both `first` and `last` to paginate a connection is not supported")
	}

	if options.First != nil && *options.First < 0 {
		return fmt.Errorf("`first` on a connection cannot be less than zero")
	}

	if options.Last != nil && *options.Last < 0 {
		return fmt.Errorf("`last` on a connection cannot be less than zero")
	}

	if options.After != nil {
		if _, err := base64.StdEncoding.DecodeString(*options.After); err != nil {
			return fmt.Errorf("invalid `after` cursor: `%s`", *options.After)
		}
	}

	if options.Before != nil {
		if _, err := base64.StdEncoding.DecodeString(*options.Before); err != nil {
			return fmt.Errorf("invalid `before` cursor: `%s`", *options.Before)
		}
	}

	return nil
}
