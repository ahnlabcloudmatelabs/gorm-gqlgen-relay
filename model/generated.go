package model

import (
	"github.com/cloudmatelabs/gorm-gqlgen-relay/filters"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/relay"
)

type PageInfo = relay.PageInfo

type StringFilter = filters.StringFilter
type IntFilter = filters.IntFilter
type FloatFilter = filters.FloatFilter
type BooleanFilter = filters.BooleanFilter
type TimeFilter = filters.TimeFilter
type UUIDFilter = filters.UUIDFilter
type IDFilter = filters.IDFilter
type MapFilter = filters.MapFilter
type MapEntry = filters.MapEntry
type MapInEntry = filters.MapInEntry
