package fields_fns

import (
	"gg/instructions"
	"gg/reflects/types"
)

func Recursive[T any](source T, field *types.Field) (instructions.Instruction, error) {

	return instructions.NOOP, nil
}
