package collections

import "cmp"

func Unique[T comparable](items []T) []T {
	panic("not implemented")
}

func UniqueBy[T any, K comparable](items []T, key func(T) K) []T {
	seen := make(map[K]struct{}, len(items))
	out := make([]T, 0, len(items))
	for _, item := range items {
		k := key(item)
		if _, ok := seen[k]; !ok {
			seen[k] = struct{}{}
			out = append(out, item)
		}
	}

	return out
}

func SortedUnique[T cmp.Ordered](items []T) []T {
	panic("not implemented")
}
