package collections

import (
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isEven(n int) bool {
	return n%2 == 0
}

func repeat(n int) []int {
	return slices.Repeat([]int{n}, n)
}

func TestFilterEmpty(t *testing.T) {
	assert.Empty(t, Filter(nil, isEven))
}

func TestFilterKeepsOrder(t *testing.T) {
	assert.Equal(t, []int{4, 2, 2}, Filter([]int{1, 4, 3, 2, 2}, isEven))
}

func TestMapEmpty(t *testing.T) {
	assert.Empty(t, Map(nil, strconv.Itoa))
}

func TestMapKeepsDuplicates(t *testing.T) {
	assert.Equal(t, []string{"1", "1", "2"}, Map([]int{1, 1, 2}, strconv.Itoa))
}

func TestFlatMapEmpty(t *testing.T) {
	assert.Empty(t, FlatMap(nil, repeat))
}

func TestFlatMapConcatenates(t *testing.T) {
	assert.Equal(t, []int{1, 2, 2, 3, 3, 3}, FlatMap([]int{1, 0, 2, 3}, repeat))
}

func TestCountFuncEmpty(t *testing.T) {
	assert.Equal(t, 0, CountFunc(nil, isEven))
}

func TestCountFunc(t *testing.T) {
	assert.Equal(t, 2, CountFunc([]int{1, 2, 3, 4}, isEven))
}
