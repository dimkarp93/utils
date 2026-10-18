package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSetEmpty(t *testing.T) {
	assert.Empty(t, ToSet[string](nil))
}

func TestToSetCollapsesDuplicates(t *testing.T) {
	assert.Equal(t, map[string]bool{"a": true, "b": true}, ToSet([]string{"a", "b", "a"}))
}

func TestSortedKeysEmpty(t *testing.T) {
	assert.Empty(t, SortedKeys[string, int](nil))
}

func TestSortedKeys(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, SortedKeys(map[string]int{"c": 1, "a": 2, "b": 3}))
}

func TestCountByEmpty(t *testing.T) {
	assert.Empty(t, CountBy(nil, pairName))
}

func TestCountByCountsDuplicates(t *testing.T) {
	items := []pair{{"a", 1}, {"b", 2}, {"a", 3}}

	assert.Equal(t, map[string]int{"a": 2, "b": 1}, CountBy(items, pairName))
}

func TestGroupByEmpty(t *testing.T) {
	assert.Empty(t, GroupBy(nil, pairName))
}

func TestGroupByKeepsOrderWithinGroup(t *testing.T) {
	items := []pair{{"a", 1}, {"b", 2}, {"a", 3}}
	want := map[string][]pair{"a": {{"a", 1}, {"a", 3}}, "b": {{"b", 2}}}

	assert.Equal(t, want, GroupBy(items, pairName))
}
