package fields_fns

import (
	"gg/reflects/types"
	"gg/stack"
)

func IgnoreStructs[T any](source T, field *types.Field) (stack.Instruction, error) {
	return stack.Next, nil
}
