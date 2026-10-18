package collections

import "cmp"

func Unique[T comparable](items []T) []T {
	panic("not implemented")
}

func UniqueBy[T any, K comparable](items []T, key func(T) K) []T {
	panic("not implemented")
}

func SortedUnique[T cmp.Ordered](items []T) []T {
	panic("not implemented")
}
