package fields_fns

import (
	"gg/reflects/assertions"
	"gg/reflects/types"
	"gg/stack"
)

func IgnorePrivate[T any](source T, field *types.Field) (stack.Instruction, error) {
	if !assertions.IsFieldExported(field.StructField) {
		return stack.Stop, nil
	}

	return stack.Next, nil
}
