package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtInRange(t *testing.T) {
	assert.Equal(t, "b", At([]string{"a", "b"}, 1))
}

func TestAtOutOfRange(t *testing.T) {
	t.Run("past the end", func(t *testing.T) {
		assert.Equal(t, "", At([]string{"a"}, 5))
	})
	t.Run("negative index", func(t *testing.T) {
		assert.Equal(t, "", At([]string{"a"}, -1))
	})
	t.Run("nil slice", func(t *testing.T) {
		assert.Equal(t, 0, At[int](nil, 0))
	})
}

func TestOrDefaultEmpty(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		assert.Equal(t, []string{"—"}, OrDefault(nil, []string{"—"}))
	})
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, []string{"—"}, OrDefault([]string{}, []string{"—"}))
	})
}

func TestOrDefaultKeepsItems(t *testing.T) {
	assert.Equal(t, []string{"a"}, OrDefault([]string{"a"}, []string{"—"}))
}
