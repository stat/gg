package executors

import (
	"gg/instructions"
	"gg/reflects/types"
)

//
// Execute Field Fns
//

func FieldFns[T any](source T, f *types.Field, fns ...types.FieldFn[T]) error {
	for _, fn := range fns {
		op, err := fn(source, f)

		if err != nil {
			return err
		}

		switch op {
		case instructions.Next:
			continue
		case instructions.Stop:
			break
		case instructions.Recurse:

		}
	}
	return nil
}
