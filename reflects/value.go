package reflects

import (
	"reflect"
)

func Value[T any](v T) reflect.Value {
	value := reflect.ValueOf(v)

	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return value
}
