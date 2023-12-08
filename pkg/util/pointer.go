package util

func P[T any](v T) *T {
	return &v
}
