package assertions

import (
	"reflect"
)

func IsFieldExported(field reflect.StructField) bool {
	return field.PkgPath == ""
}
