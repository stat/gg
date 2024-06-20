package fields_fns

import (
	"gg/instructions"
	"gg/reflects/assertions"
	"gg/reflects/types"
)

func IgnorePrivate[T any](source T, field *types.Field) (instructions.Instruction, error) {
	if !assertions.IsFieldExported(field.StructField) {
		return instructions.Stop, nil
	}

	return instructions.Next, nil
}
