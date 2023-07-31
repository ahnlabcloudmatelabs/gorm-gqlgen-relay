package model

import (
	"github.com/cloudmatelabs/gorm-gqlgen-relay/filters"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/relay"
)

type PageInfo = relay.PageInfo

type IDFilter = filters.IDFilter
type IntFilter = filters.IntFilter
type MapFilter = filters.MapFilter
type StringFilter = filters.StringFilter
type TimeFilter = filters.TimeFilter
type UUIDFilter = filters.UUIDFilter
