package types

import (
	"reflect"

	"gg/stack"
)

// Field

type Field struct {
	reflect.StructField
	reflect.Type
}

// FieldFn

type FieldFn[T any] func(T, *Field) (stack.Instruction, error)
