package reflects

type Tag[V any] struct {
	Field     any
	FieldName string
	Value     V
}

type TagFn[T any, V any] func(T, V) error

func Tags[T any, V any](t T, v []V, fns ...TagFn[T, V]) error {
	return nil
}
