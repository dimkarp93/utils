# Использование

*Языки: [English](usage.md) · **Русский***

```go
import "github.com/dimkarp93/utils/pkg/collections"
```

## Цепочки

```go
invalid := collections.Map(
	collections.Filter(collections.SortedKeys(envToName), isInvalidEnvName),
	strconv.Quote,
)

env := collections.Filter(environ, func(kv string) bool { return envKey(kv) != passwordEnv })

titles := collections.Unique(collections.Map(matches, submatchName))
```

## Группировка и подсчёт

```go
byStore := collections.GroupBy(sectionUsages(cfg), usageKeyStore)
for ks, group := range byStore {
	out[ks] = collections.SortedUnique(collections.FlatMap(group, usageTitles))
}

localCount := collections.CountBy(locals, repoBaseName)

failed := collections.CountFunc(results, func(r fetchResult) bool { return r.Err != nil })
```

## Уникальность и множества

```go
envPairs := collections.Unique(collections.FlatMap(group, usageEnvPairs))
hasEnv := collections.ToSet(collections.Map(envPairs, resolve.MappingName))
channels := collections.UniqueBy(collections.FlatMap(group, usageChannels), resolve.MappingName)

targets := collections.Unique(collections.Filter(paths, func(p string) bool { return p != "" }))
```

## Безопасный доступ и значения по умолчанию

```go
text := Cell(collections.At(cols[0], i), width)

rem := collections.OrDefault(Items(r.Remotes, width), []string{"—"})
```

## Детерминированный обход

```go
for _, name := range collections.SortedKeys(cfg.Sections) {
}

merged := collections.SortedUnique(slices.Concat(current, added))
```

## Если хватает стандартной библиотеки — берите её

`slices.Concat`, `slices.ContainsFunc`, `slices.IndexFunc`, `slices.Repeat`, `slices.DeleteFunc`,
`maps.Clone`, `maps.Copy`, `cmp.Or` и встроенные `max` / `min` отдельных хелперов не требуют.
