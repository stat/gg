package types

import (
	"reflect"

	"gg/instructions"
)

// Field

type Field struct {
	reflect.StructField
	reflect.Type
}

// FieldFn

type FieldFn[T any] func(T, *Field) (instructions.Instruction, error)
