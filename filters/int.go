package filters

type IntFilter struct {
	Eq        *int   `json:"eq,omitempty"`
	Ne        *int   `json:"ne,omitempty"`
	In        *[]int `json:"in,omitempty"`
	Nin       *[]int `json:"nIn,omitempty"`
	Lt        *int   `json:"lt,omitempty"`
	Lte       *int   `json:"lte,omitempty"`
	Gt        *int   `json:"gt,omitempty"`
	Gte       *int   `json:"gte,omitempty"`
	IsNull    bool   `json:"isNull,omitempty"`
	IsNotNull bool   `json:"isNotNull,omitempty"`
}

func (f *Filter) Int(filter interface{}, where bool) {
	filterData := parseFilter(filter)

	f.Equal("eq", filterData["eq"], where)
	f.NotEqual("ne", filterData["ne"], where)
	f.In("in", filterData["in"], where)
	f.NotIn("nin", filterData["nin"], where)
	f.LessThan("lt", filterData["lt"], where)
	f.LessThanOrEqual("lte", filterData["lte"], where)
	f.GreaterThan("gt", filterData["gt"], where)
	f.GreaterThanOrEqual("gte", filterData["gte"], where)
	f.IsNull("isnull", filterData["isnull"], where)
	f.IsNotNull("isnotnull", filterData["isnotnull"], where)
}
