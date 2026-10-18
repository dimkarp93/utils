# utils

*Языки: [English](README.md) · **Русский***

Общая внутренняя библиотека небольших функций на Go без состояния. Пока в ней дженерик-операции
над коллекциями в `pkg/collections`.

Модуль: `github.com/dimkarp93/utils`. Go 1.27.2. В коде библиотеки нет зависимостей;
[testify](https://github.com/stretchr/testify) используется только в тестах.

## Установка

Репозиторий приватный, поэтому `go` нужно запретить ходить за ним в публичный прокси:

```sh
go env -w GOPRIVATE='github.com/dimkarp93/*'
go get github.com/dimkarp93/utils@latest
```

```go
import "github.com/dimkarp93/utils/pkg/collections"
```

## collections

| Функция | Сигнатура | Поведение |
|---|---|---|
| `Filter` | `Filter[T any](items []T, keep func(T) bool) []T` | оставляет подходящие элементы, порядок сохраняется |
| `Map` | `Map[T, R any](items []T, fn func(T) R) []R` | преобразует каждый элемент, дубликаты сохраняются |
| `FlatMap` | `FlatMap[T, R any](items []T, fn func(T) []R) []R` | преобразует в срезы и склеивает их |
| `Unique` | `Unique[T comparable](items []T) []T` | убирает дубликаты, остаётся первое вхождение и порядок |
| `UniqueBy` | `UniqueBy[T any, K comparable](items []T, key func(T) K) []T` | как `Unique`, но сравнение по ключу |
| `SortedUnique` | `SortedUnique[T cmp.Ordered](items []T) []T` | сортирует и убирает дубликаты, входной срез не меняется |
| `GroupBy` | `GroupBy[T any, K comparable](items []T, key func(T) K) map[K][]T` | группирует по ключу, порядок внутри группы сохраняется |
| `CountBy` | `CountBy[T any, K comparable](items []T, key func(T) K) map[K]int` | число вхождений по ключу |
| `CountFunc` | `CountFunc[T any](items []T, keep func(T) bool) int` | число подходящих элементов |
| `ToSet` | `ToSet[T comparable](items []T) map[T]bool` | множество для проверок `set[x]` |
| `SortedKeys` | `SortedKeys[K cmp.Ordered, V any](m map[K]V) []K` | ключи map по возрастанию |
| `At` | `At[T any](items []T, i int) T` | `items[i]`, нулевое значение при выходе за границы |
| `OrDefault` | `OrDefault[T any](items []T, def []T) []T` | `def`, если `items` пуст |

Примеры — в [usage.ru.md](usage.ru.md); правила для разработчиков и агентов — в
[AGENTS.md](AGENTS.md).

## Сборка

```sh
make check        # go vet + go test
make cover        # покрытие
make bump-patch   # 1.2.3 -> 1.2.4, тег, push в origin
```

Версия лежит в `versions.txt`; тег `vX.Y.Z` запускает workflow релиза.

## Лицензия

[MIT](LICENSE)
