package filters

import (
	"github.com/cloudmatelabs/gorm-gqlgen-relay/filters/mapfilter"
	"gorm.io/datatypes"
)

type MapFilter struct {
	Eq  interface{} `json:"eq,omitempty"`
	Ne  interface{} `json:"ne,omitempty"`
	In  *[]string   `json:"in,omitempty"`
	Nin *[]string   `json:"nIn,omitempty"`
}

func (f *Filter) Map(field string, filter *MapFilter, where bool) {
	if filter == nil {
		return
	}

	f.MapEqual(field, filter.Eq, where)
	f.MapNotEqual(field, filter.Ne, where)
	f.MapIn(field, filter.In, where)
	f.MapNotIn(field, filter.Nin, where)
}

func (f *Filter) MapEqual(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	if _, ok := value.([]interface{}); ok {
		f.MapArrayEqual(field, value, where)
		return
	}

	dbType := f.DB.Dialector.Name()

	if dbType == "sqlserver" {
		f.DB = mapfilter.MapEqualSqlServer(f.DB, field, value.(map[string]interface{}), where)
		return
	}

	if where {
		f.DB = f.DB.Where(datatypes.JSONQuery(field).Equals(value))
		return
	}

	f.DB = f.DB.Or(datatypes.JSONQuery(field).Equals(value))
}

func (f *Filter) MapArrayEqual(field string, value interface{}, where bool) {
	dbType := f.DB.Dialector.Name()

	if dbType == "sqlserver" {
		f.DB = mapfilter.MapArrayEqualSqlServer(f.DB, field, value.([]map[string]interface{}), where)
		return
	}

	if where {
		f.DB = f.DB.Where(datatypes.JSONArrayQuery(field).Contains(value))
		return
	}

	f.DB = f.DB.Or(datatypes.JSONArrayQuery(field).Contains(value))
}

func (f *Filter) MapNotEqual(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	dbType := f.DB.Dialector.Name()

	if dbType == "sqlserver" {
		f.DB = mapfilter.MapNotEqualSqlServer(f.DB, field, value.(map[string]interface{}), where)
		return
	}

	if where {
		f.DB = f.DB.Not(datatypes.JSONQuery(field).Equals(value))
		return
	}

	f.DB = f.DB.Or(field+" != ?", value)
}

func (f *Filter) MapArrayNotEqual(field string, value interface{}, where bool) {
	dbType := f.DB.Dialector.Name()

	if dbType == "sqlserver" {
		f.DB = mapfilter.MapArrayNotEqualSqlServer(f.DB, field, value.([]map[string]interface{}), where)
		return
	}

	if where {
		f.DB = f.DB.Not(datatypes.JSONArrayQuery(field).Contains(value))
		return
	}

	f.DB = f.DB.Or(datatypes.JSONArrayQuery(field).Contains(value))
}

func (f *Filter) MapIn(field string, value *[]string, where bool) {
	if value == nil {
		return
	}

	dbType := f.DB.Dialector.Name()

	if dbType == "sqlserver" {
		f.DB = mapfilter.MapInSqlServer(f.DB, field, *value, where)
		return
	}

	if where {
		f.DB = f.DB.Where(datatypes.JSONQuery(field).HasKey(*value...))
		return
	}

	f.DB = f.DB.Or(datatypes.JSONQuery(field).HasKey(*value...))
}

func (f *Filter) MapNotIn(field string, value *[]string, where bool) {
	if value == nil {
		return
	}

	dbType := f.DB.Dialector.Name()

	if dbType == "sqlserver" {
		f.DB = mapfilter.MapNotInSqlServer(f.DB, field, *value, where)
		return
	}

	if where {
		f.DB = f.DB.Not(datatypes.JSONQuery(field).HasKey(*value...))
		return
	}

	f.DB = f.DB.Or(datatypes.JSONQuery(field).HasKey(*value...))
}
