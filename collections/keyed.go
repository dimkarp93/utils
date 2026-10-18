package collections

import (
	"cmp"
)

func ToSet[T comparable](items []T) map[T]bool {
	set := make(map[T]bool, len(items))
	for _, item := range items {
		set[item] = true
	}

	return set
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
