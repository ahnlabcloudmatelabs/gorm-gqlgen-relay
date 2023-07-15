package filters

import (
	"fmt"

	"gorm.io/datatypes"
)

type MapFilter struct {
	Eq  *map[string]interface{} `json:"eq,omitempty"`
	Ne  *map[string]interface{} `json:"ne,omitempty"`
	In  *[]string               `json:"in,omitempty"`
	Nin *[]string               `json:"nIn,omitempty"`
}

func (f *Filter) Map(filter interface{}, where bool) {
	filterData := parseFilter(filter)

	f.MapEqual("eq", filterData["eq"], where)
	f.MapNotEqual("ne", filterData["ne"], where)
	f.MapIn("contains", filterData["contains"], where)
	f.MapNotIn("ncontains", filterData["ncontains"], where)
}

func (f *Filter) MapEqual(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	dbType := f.DB.Dialector.Name()

	if dbType == "sqlserver" {
		f.mapEqualSqlServer(field, value, where)
	}

	if where {
		f.DB = f.DB.Where(datatypes.JSONQuery(field).Equals(value))
		return
	}

	f.DB = f.DB.Or(datatypes.JSONQuery(field).Equals(value))
}

func (f *Filter) mapEqualSqlServer(field string, value interface{}, where bool) {
	if where {
		f.DB = f.DB.Where(fmt.Sprintf("JSON_VALUE(%s, '$.%s') = ?", field, value))
		return
	}

	f.DB = f.DB.Or(fmt.Sprintf("JSON_VALUE(%s, '$.%s') = ?", field, value))
}

func (f *Filter) MapNotEqual(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	dbType := f.DB.Dialector.Name()

	if dbType == "sqlserver" {
		f.mapNotEqualSqlServer(field, value, where)
	}

	if where {
		f.DB = f.DB.Not(datatypes.JSONQuery(field).Equals(value))
		return
	}

	f.DB = f.DB.Or(field+" != ?", value)
}

func (f *Filter) mapNotEqualSqlServer(field string, value interface{}, where bool) {
	if where {
		f.DB = f.DB.Where(fmt.Sprintf("JSON_VALUE(%s, '$.%s') != ?", field, value))
		return
	}

	f.DB = f.DB.Or(fmt.Sprintf("JSON_VALUE(%s, '$.%s') != ?", field, value))
}

func (f *Filter) MapIn(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	dbType := f.DB.Dialector.Name()

	if dbType == "sqlserver" {
		f.mapContainsSqlServer(field, value.([]string), where)
		return
	}

	if where {
		f.DB = f.DB.Where(datatypes.JSONQuery(field).HasKey(value.([]string)...))
		return
	}

	f.DB = f.DB.Or(datatypes.JSONQuery(field).HasKey(value.([]string)...))
}

func (f *Filter) mapContainsSqlServer(field string, values []string, where bool) {
	if where {
		for _, value := range values {
			f.DB = f.DB.Where(fmt.Sprintf("JSON_VALUE(%s, '$.%s') IS NOT NULL", field, value))
		}
		return
	}

	for _, value := range values {
		f.DB = f.DB.Or(fmt.Sprintf("JSON_VALUE(%s, '$.%s') IS NOT NULL", field, value))
	}
}

func (f *Filter) MapNotIn(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	dbType := f.DB.Dialector.Name()

	if dbType == "sqlserver" {
		f.mapNotContainsSqlServer(field, value.([]string), where)
		return
	}

	if where {
		f.DB = f.DB.Not(datatypes.JSONQuery(field).HasKey(value.([]string)...))
		return
	}

	f.DB = f.DB.Or(datatypes.JSONQuery(field).HasKey(value.([]string)...))
}

func (f *Filter) mapNotContainsSqlServer(field string, values []string, where bool) {
	if where {
		for _, value := range values {
			f.DB = f.DB.Where(fmt.Sprintf("JSON_VALUE(%s, '$.%s') IS NULL", field, value))
		}
		return
	}

	for _, value := range values {
		f.DB = f.DB.Or(fmt.Sprintf("JSON_VALUE(%s, '$.%s') IS NULL", field, value))
	}
}
