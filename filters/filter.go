package filters

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type Filter struct {
	DB *gorm.DB
}

func parseFilter(filter interface{}) (data map[string]interface{}) {
	byteData, _ := json.Marshal(filter)
	json.Unmarshal(byteData, &data)
	return
}

func (f *Filter) Equal(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	if where {
		f.DB = f.DB.Where(field+" = ?", value)
		return
	}

	f.DB = f.DB.Or(field+" = ?", value)
}

func (f *Filter) NotEqual(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	if where {
		f.DB = f.DB.Where(field+" != ?", value)
		return
	}

	f.DB = f.DB.Or(field+" != ?", value)
}

func (f *Filter) In(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	if where {
		f.DB = f.DB.Where(field+" IN ?", value)
		return
	}

	f.DB = f.DB.Or(field+" IN ?", value)
}

func (f *Filter) NotIn(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	if where {
		f.DB = f.DB.Where(field+" NOT IN ?", value)
		return
	}

	f.DB = f.DB.Or(field+" NOT IN ?", value)
}

func (f *Filter) Contains(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	_value := fmt.Sprintf("%%%v%%", value)

	if where {
		f.DB = f.DB.Where(field+" LIKE ?", _value)
		return
	}

	f.DB = f.DB.Or(field+" LIKE ?", _value)
}

func (f *Filter) NotContains(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	_value := fmt.Sprintf("%%%v%%", value)

	if where {
		f.DB = f.DB.Where(field+" NOT LIKE ?", _value)
		return
	}

	f.DB = f.DB.Or(field+" NOT LIKE ?", _value)
}

func (f *Filter) StartsWith(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	_value := fmt.Sprintf("%v%%", value)

	if where {
		f.DB = f.DB.Where(field+" LIKE ?", _value)
		return
	}

	f.DB = f.DB.Or(field+" LIKE ?", _value)
}

func (f *Filter) NotStartsWith(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	_value := fmt.Sprintf("%v%%", value)

	if where {
		f.DB = f.DB.Where(field+" NOT LIKE ?", _value)
		return
	}

	f.DB = f.DB.Or(field+" NOT LIKE ?", _value)
}

func (f *Filter) EndsWith(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	_value := fmt.Sprintf("%%%v", value)

	if where {
		f.DB = f.DB.Where(field+" LIKE ?", _value)
		return
	}

	f.DB = f.DB.Or(field+" LIKE ?", _value)
}

func (f *Filter) NotEndsWith(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	_value := fmt.Sprintf("%%%v", value)

	if where {
		f.DB = f.DB.Where(field+" NOT LIKE ?", _value)
		return
	}

	f.DB = f.DB.Or(field+" NOT LIKE ?", _value)
}

func (f *Filter) LessThan(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	if where {
		f.DB = f.DB.Where(field+" < ?", value)
		return
	}

	f.DB = f.DB.Or(field+" < ?", value)
}

func (f *Filter) LessThanOrEqual(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	if where {
		f.DB = f.DB.Where(field+" <= ?", value)
		return
	}

	f.DB = f.DB.Or(field+" <= ?", value)
}

func (f *Filter) GreaterThan(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	if where {
		f.DB = f.DB.Where(field+" > ?", value)
		return
	}

	f.DB = f.DB.Or(field+" > ?", value)
}

func (f *Filter) GreaterThanOrEqual(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	if where {
		f.DB = f.DB.Where(field+" >= ?", value)
		return
	}

	f.DB = f.DB.Or(field+" >= ?", value)
}

func (f *Filter) Between(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	if where {
		f.DB = f.DB.Where(field+" BETWEEN ?", value)
		return
	}

	f.DB = f.DB.Or(field+" BETWEEN ?", value)
}

func (f *Filter) NotBetween(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	if where {
		f.DB = f.DB.Where(field+" NOT BETWEEN ?", value)
		return
	}

	f.DB = f.DB.Or(field+" NOT BETWEEN ?", value)
}

func (f *Filter) IsNull(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	if where {
		f.DB = f.DB.Where(field + " IS NULL")
		return
	}

	f.DB = f.DB.Or(field + " IS NULL")
}

func (f *Filter) IsNotNull(field string, value interface{}, where bool) {
	if value == nil {
		return
	}

	if where {
		f.DB = f.DB.Where(field + " IS NOT NULL")
		return
	}

	f.DB = f.DB.Or(field + " IS NOT NULL")
}
