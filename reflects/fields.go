package reflects

import (
	"gg/instructions"
	"gg/reflects/types"
)

//
// Types
//

//
// Fields
//

func Fields[T any](source T, fns ...types.FieldFn[T]) error {
	v := Value(source)
	t := v.Type()
	n := t.NumField()

	for i := 0; i < n; i++ {
		f := t.Field(i)

		field := &types.Field{
			StructField: f,
			Type:        t,
		}

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

func ExecuteFieldFns[T any](source T, f *types.Field, fns ...types.FieldFn[T]) error {
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
