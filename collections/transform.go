package collections

func Filter[T any](items []T, keep func(T) bool) []T {
	panic("not implemented")
}

func Map[T, R any](items []T, fn func(T) R) []R {
	panic("not implemented")
}

func FlatMap[T, R any](items []T, fn func(T) []R) []R {
	panic("not implemented")
}

func CountFunc[T any](items []T, keep func(T) bool) int {
	panic("not implemented")
}
