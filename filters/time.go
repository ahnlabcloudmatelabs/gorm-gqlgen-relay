package filters

import "time"

type TimeFilter struct {
	Eq        *time.Time   `json:"eq,omitempty"`
	Ne        *time.Time   `json:"ne,omitempty"`
	In        *[]time.Time `json:"in,omitempty"`
	Nin       *[]time.Time `json:"nIn,omitempty"`
	Lt        *time.Time   `json:"lt,omitempty"`
	Lte       *time.Time   `json:"lte,omitempty"`
	Gt        *time.Time   `json:"gt,omitempty"`
	Gte       *time.Time   `json:"gte,omitempty"`
	IsNull    *bool        `json:"isNull,omitempty"`
	IsNotNull *bool        `json:"isNotNull,omitempty"`
}

func (f *Filter) Time(filter interface{}, where bool) {
	filterData := parseFilter(filter)

	f.Equal("eq", filterData["eq"], where)
	f.NotEqual("ne", filterData["ne"], where)
	f.In("in", filterData["in"], where)
	f.NotIn("nin", filterData["nin"], where)
	f.GreaterThan("gt", filterData["gt"], where)
	f.GreaterThanOrEqual("gte", filterData["gte"], where)
	f.LessThan("lt", filterData["lt"], where)
	f.LessThanOrEqual("lte", filterData["lte"], where)
	f.IsNull("isnull", filterData["isnull"], where)
	f.IsNotNull("isnotnull", filterData["isnotnull"], where)
}
