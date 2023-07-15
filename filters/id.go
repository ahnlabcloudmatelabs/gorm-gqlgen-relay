package filters

type IDFilter struct {
	Eq          interface{}   `json:"eq,omitempty"`
	Ne          interface{}   `json:"ne,omitempty"`
	Lt          interface{}   `json:"lt,omitempty"`
	Lte         interface{}   `json:"lte,omitempty"`
	Gt          interface{}   `json:"gt,omitempty"`
	Gte         interface{}   `json:"gte,omitempty"`
	In          []interface{} `json:"in,omitempty"`
	Nin         []interface{} `json:"nin,omitempty"`
	Contains    interface{}   `json:"contains,omitempty"`
	NContains   interface{}   `json:"nContains,omitempty"`
	StartsWith  interface{}   `json:"startsWith,omitempty"`
	NStartsWith interface{}   `json:"nStartsWith,omitempty"`
	EndsWith    interface{}   `json:"endsWith,omitempty"`
	NEndsWith   interface{}   `json:"nEndsWith,omitempty"`
	IsNull      bool          `json:"isNull,omitempty"`
	IsNotNull   bool          `json:"isNotNull,omitempty"`
}

func (f *Filter) ID(filter interface{}, where bool) {
	filterData := parseFilter(filter)

	f.Equal("eq", filterData["eq"], where)
	f.NotEqual("ne", filterData["ne"], where)
	f.LessThan("lt", filterData["lt"], where)
	f.LessThanOrEqual("lte", filterData["lte"], where)
	f.GreaterThan("gt", filterData["gt"], where)
	f.GreaterThanOrEqual("gte", filterData["gte"], where)
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
