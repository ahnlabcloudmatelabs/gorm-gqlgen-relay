package filters

type StringFilter struct {
	Eq          *string   `json:"eq,omitempty"`
	Ne          *string   `json:"ne,omitempty"`
	In          *[]string `json:"in,omitempty"`
	Nin         *[]string `json:"nIn,omitempty"`
	Contains    *string   `json:"contains,omitempty"`
	NContains   *string   `json:"nContains,omitempty"`
	StartsWith  *string   `json:"startsWith,omitempty"`
	NStartsWith *string   `json:"nStartsWith,omitempty"`
	EndsWith    *string   `json:"endsWith,omitempty"`
	NEndsWith   *string   `json:"nEndsWith,omitempty"`
	IsNull      *bool     `json:"isNull,omitempty"`
	IsNotNull   *bool     `json:"isNotNull,omitempty"`
}

func (f *Filter) String(filter interface{}, where bool) {
	filterData := parseFilter(filter)

	f.Equal("eq", filterData["eq"], where)
	f.NotEqual("ne", filterData["ne"], where)
	f.In("in", filterData["in"], where)
	f.NotIn("nin", filterData["nin"], where)
	f.Contains("contains", filterData["contains"], where)
	f.NotContains("ncontains", filterData["ncontains"], where)
	f.StartsWith("startswith", filterData["startswith"], where)
	f.NotStartsWith("nstartswith", filterData["nstartswith"], where)
	f.EndsWith("endswith", filterData["endswith"], where)
	f.NotEndsWith("nendswith", filterData["nendswith"], where)
	f.IsNull("isnull", filterData["isnull"], where)
	f.IsNotNull("isnotnull", filterData["isnotnull"], where)
}
