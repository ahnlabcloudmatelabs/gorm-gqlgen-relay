package filters

import "github.com/google/uuid"

type UUIDFilter struct {
	Eq        *uuid.UUID   `json:"eq,omitempty"`
	Ne        *uuid.UUID   `json:"ne,omitempty"`
	In        *[]uuid.UUID `json:"in,omitempty"`
	Nin       *[]uuid.UUID `json:"nIn,omitempty"`
	IsNull    *bool        `json:"isNull,omitempty"`
	IsNotNull *bool        `json:"isNotNull,omitempty"`
}

func (f *Filter) UUID(filter interface{}, where bool) {
	filterData := parseFilter(filter)

	f.Equal("eq", filterData["eq"], where)
	f.NotEqual("ne", filterData["ne"], where)
	f.In("in", filterData["in"], where)
	f.NotIn("nin", filterData["nin"], where)
	f.IsNull("isnull", filterData["isnull"], where)
	f.IsNotNull("isnotnull", filterData["isnotnull"], where)
}
