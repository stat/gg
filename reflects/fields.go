package reflects

import (
	"gg/reflects/executors"
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
	n := v.NumField()

	for i := 0; i < n; i++ {
		f := t.Field(i)

		field := &types.Field{
			StructField: f,
			Type:        t,
		}

		err := executors.FieldFns(source, field, fns...)

		if err != nil {
			return err
		}
	}

	return nil
}
