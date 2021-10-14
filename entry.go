package entry

import (
	"entry/array"
	"entry/single"
	"reflect"
)

func IsContains(what interface{}, where interface{}) bool {
	s := reflect.ValueOf(what)
	if s.Kind() == reflect.Slice || s.Kind() == reflect.Array {
		return array.SwitchType(what, where)
	}
	return single.SwitchType(what, where)
}
