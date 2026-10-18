package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type pair struct {
	name  string
	value int
}

func pairName(p pair) string {
	return p.name
}

func TestUniqueEmpty(t *testing.T) {
	assert.Empty(t, Unique[string](nil))
}

func TestUniqueKeepsFirstOccurrence(t *testing.T) {
	assert.Equal(t, []string{"b", "a", "c"}, Unique([]string{"b", "a", "b", "c", "a"}))
}

func TestUniqueByEmpty(t *testing.T) {
	assert.Empty(t, UniqueBy(nil, pairName))
}

func TestUniqueByKeepsFirstForKey(t *testing.T) {
	items := []pair{{"a", 1}, {"b", 2}, {"a", 3}}

	assert.Equal(t, []pair{{"a", 1}, {"b", 2}}, UniqueBy(items, pairName))
}

func TestSortedUniqueEmpty(t *testing.T) {
	assert.Empty(t, SortedUnique[string](nil))
}

func TestSortedUniqueDropsDuplicates(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, SortedUnique([]string{"c", "a", "b", "a", "c"}))
}

func TestSortedUniqueKeepsInput(t *testing.T) {
	items := []string{"c", "a", "c"}

	SortedUnique(items)

	assert.Equal(t, []string{"c", "a", "c"}, items)
}
