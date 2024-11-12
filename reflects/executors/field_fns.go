package executors

import (
	"gg/reflects/types"
	"gg/stack"
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
		case stack.Next:
			continue
		case stack.Stop:
			break
		case stack.Recurse:

		}
	}
	return nil
}
