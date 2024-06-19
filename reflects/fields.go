package reflects

import (
	"reflect"

	"gg/instructions"
)

//
// Types
//

type Field struct {
	reflect.StructField
	reflect.Type
}

type FieldFn[T any] func(T, *Field) (instructions.Instruction, error)

//
// Fields
//

func Fields[T any](source T, fns ...FieldFn[T]) error {
	v := Value(source)
	t := v.Type()
	n := t.NumField()

	for i := 0; i < n; i++ {
		f := t.Field(i)
		field := &Field{f, t}

		err := ExecuteFieldFns(source, field, fns...)

		if err != nil {
			return err
		}
	}

	return nil
}

//
// Execute Field Fns
//

func ExecuteFieldFns[T any](source T, f *Field, fns ...FieldFn[T]) error {
	for _, fn := range fns {
		op, err := fn(source, f)

		if err != nil {
			return err
		}

		switch op {
		case instructions.Continue:
			continue
		case instructions.Skip:
			break
		}
	}
	return nil
}
