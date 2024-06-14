package reflects

type FieldFn[T any] func(T) error

func Fields[T any](v T, fns ...FieldFn[T]) error {
	return nil
}
