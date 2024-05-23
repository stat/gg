package utils

func ByRef[T any](v T) *T {
	return &v
}
