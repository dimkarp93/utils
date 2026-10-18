package collections

func Filter[T any](items []T, keep func(T) bool) []T {
	out := make([]T, 0, len(items))
	for _, item := range items {
		if keep(item) {
			out = append(out, item)
		}
	}

	return out
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
