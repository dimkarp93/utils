# utils

*Languages: **English** · [Русский](README.ru.md)*

Shared internal Go library of small, stateless helpers. For now it holds generic collection
operations in `pkg/collections`.

Module: `github.com/dimkarp93/utils`. Go 1.27.2. No dependencies in the library code;
[testify](https://github.com/stretchr/testify) is used by the tests only.

## Install

The repository is private, so tell `go` not to use the public proxy for it:

```sh
go env -w GOPRIVATE='github.com/dimkarp93/*'
go get github.com/dimkarp93/utils@latest
```

```go
import "github.com/dimkarp93/utils/pkg/collections"
```

## collections

| Function | Signature | Behaviour |
|---|---|---|
| `Filter` | `Filter[T any](items []T, keep func(T) bool) []T` | keeps matching items, order preserved |
| `Map` | `Map[T, R any](items []T, fn func(T) R) []R` | maps every item, duplicates preserved |
| `FlatMap` | `FlatMap[T, R any](items []T, fn func(T) []R) []R` | maps to slices and concatenates them |
| `Unique` | `Unique[T comparable](items []T) []T` | drops duplicates, keeps the first occurrence and the order |
| `UniqueBy` | `UniqueBy[T any, K comparable](items []T, key func(T) K) []T` | like `Unique`, compares by key |
| `SortedUnique` | `SortedUnique[T cmp.Ordered](items []T) []T` | sorts and drops duplicates, the input is not modified |
| `GroupBy` | `GroupBy[T any, K comparable](items []T, key func(T) K) map[K][]T` | groups by key, order inside a group preserved |
| `CountBy` | `CountBy[T any, K comparable](items []T, key func(T) K) map[K]int` | occurrences per key |
| `CountFunc` | `CountFunc[T any](items []T, keep func(T) bool) int` | number of matching items |
| `ToSet` | `ToSet[T comparable](items []T) map[T]bool` | builds a set for `set[x]` lookups |
| `SortedKeys` | `SortedKeys[K cmp.Ordered, V any](m map[K]V) []K` | map keys in ascending order |
| `At` | `At[T any](items []T, i int) T` | `items[i]`, zero value when `i` is out of range |
| `OrDefault` | `OrDefault[T any](items []T, def []T) []T` | `def` when `items` is empty |

Examples are in [usage.md](usage.md); the rules for contributors and agents are in
[AGENTS.md](AGENTS.md).

## Build

```sh
make check        # go vet + go test
make cover        # coverage
make bump-patch   # 1.2.3 -> 1.2.4, tag, push to origin
```

The version lives in `versions.txt`; a `vX.Y.Z` tag triggers the release workflow.

## License

[Apache-2.0](LICENSE)
