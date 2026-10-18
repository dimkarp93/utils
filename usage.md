# Usage

*Languages: **English** · [Русский](usage.ru.md)*

```go
import "github.com/dimkarp93/utils/pkg/collections"
```

## Chains

```go
invalid := collections.Map(
	collections.Filter(collections.SortedKeys(envToName), isInvalidEnvName),
	strconv.Quote,
)

env := collections.Filter(environ, func(kv string) bool { return envKey(kv) != passwordEnv })

titles := collections.Unique(collections.Map(matches, submatchName))
```

## Grouping and counting

```go
byStore := collections.GroupBy(sectionUsages(cfg), usageKeyStore)
for ks, group := range byStore {
	out[ks] = collections.SortedUnique(collections.FlatMap(group, usageTitles))
}

localCount := collections.CountBy(locals, repoBaseName)

failed := collections.CountFunc(results, func(r fetchResult) bool { return r.Err != nil })
```

## Uniqueness and sets

```go
envPairs := collections.Unique(collections.FlatMap(group, usageEnvPairs))
hasEnv := collections.ToSet(collections.Map(envPairs, resolve.MappingName))
channels := collections.UniqueBy(collections.FlatMap(group, usageChannels), resolve.MappingName)

targets := collections.Unique(collections.Filter(paths, func(p string) bool { return p != "" }))
```

## Safe access and defaults

```go
text := Cell(collections.At(cols[0], i), width)

rem := collections.OrDefault(Items(r.Remotes, width), []string{"—"})
```

## Deterministic iteration

```go
for _, name := range collections.SortedKeys(cfg.Sections) {
}

merged := collections.SortedUnique(slices.Concat(current, added))
```

## Use the standard library when it is enough

`slices.Concat`, `slices.ContainsFunc`, `slices.IndexFunc`, `slices.Repeat`, `slices.DeleteFunc`,
`maps.Clone`, `maps.Copy`, `cmp.Or` and the built-in `max` / `min` need no helper here.
