package collections

import (
	"cmp"
)

func ToSet[T comparable](items []T) map[T]bool {
	panic("not implemented")
}

func SortedKeys[K cmp.Ordered, V any](m map[K]V) []K {
	panic("not implemented")
}

func CountBy[T any, K comparable](items []T, key func(T) K) map[K]int {
	panic("not implemented")
}

func GroupBy[T any, K comparable](items []T, key func(T) K) map[K][]T {
	panic("not implemented")
}
