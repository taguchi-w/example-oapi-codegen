package util

func P[T any](v T) *T {
	return &v
}

func PorNil[T comparable](v T, zero T) *T {
	if v == zero {
		return nil
	}
	return &v
}
