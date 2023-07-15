package filters

type TimeFilter struct {
	Eq        string   `json:"eq,omitempty"`
	Ne        string   `json:"ne,omitempty"`
	In        []string `json:"in,omitempty"`
	Nin       []string `json:"nIn,omitempty"`
	Lt        string   `json:"lt,omitempty"`
	Lte       string   `json:"lte,omitempty"`
	Gt        string   `json:"gt,omitempty"`
	Gte       string   `json:"gte,omitempty"`
	IsNull    bool     `json:"isNull,omitempty"`
	IsNotNull bool     `json:"isNotNull,omitempty"`
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
