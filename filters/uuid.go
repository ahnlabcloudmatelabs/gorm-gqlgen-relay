package filters

type UUIDFilter struct {
	Eq        string   `json:"eq,omitempty"`
	Ne        string   `json:"ne,omitempty"`
	In        []string `json:"in,omitempty"`
	Nin       []string `json:"nIn,omitempty"`
	IsNull    bool     `json:"isNull,omitempty"`
	IsNotNull bool     `json:"isNotNull,omitempty"`
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
