package utils

func ToPtr[T any](val T) *T {
	result := new(T)
	*result = val
	return result
}

type Range[T any] struct {
	From, To *T
}
