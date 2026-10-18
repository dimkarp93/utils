# AGENTS

Import `github.com/dimkarp93/utils/pkg/collections` (set `GOPRIVATE=github.com/dimkarp93/*`) and use its generic helpers instead of hand-written loops or private one-off helpers in the consuming project. This library holds only stateless collection operations; the library code has no dependencies beyond the standard library, and testify is allowed in tests only. Every new helper must be generic, free of state and I/O, covered by a test for empty input and its key property, and written without comments.
