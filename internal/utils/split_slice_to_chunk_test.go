package utils_test

import (
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/ozonva/ova-promise-api/internal/utils"
)

//nolint:dupl // tests
func TestSplitToChunkStrings(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		split := utils.SplitSliceToChunkStrings([]string{}, 2)

		assert.Equal(t, nil, split)
	})

	t.Run("one item", func(t *testing.T) {
		split := utils.SplitSliceToChunkStrings([]string{"hello"}, 2)

		assert.Equal(t, [][]string{{"hello"}}, split)
	})

	t.Run("multiple items", func(t *testing.T) {
		sl := []string{"a", "b", "c", "d", "e"}

		split := utils.SplitSliceToChunkStrings(sl, 2)

		assert.Equal(t, [][]string{{"a", "b"}, {"c", "d"}, {"e"}}, split)
	})

	t.Run("multiple items and zero chunkSize", func(t *testing.T) {
		sl := []string{"a", "b", "c", "d", "e"}

		split := utils.SplitSliceToChunkStrings(sl, 0)

		assert.Equal(t, [][]string{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}}, split)
	})
}

//nolint:dupl // tests
func TestSplitSliceToChunkIntegers(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		split := utils.SplitSliceToChunkIntegers([]int{}, 2)

		assert.Equal(t, nil, split)
	})

	t.Run("one item", func(t *testing.T) {
		split := utils.SplitSliceToChunkIntegers([]int{42}, 2)

		assert.Equal(t, [][]int{{42}}, split)
	})

	t.Run("multiple items", func(t *testing.T) {
		sl := []int{0, 1, 2, 3, 4}

		split := utils.SplitSliceToChunkIntegers(sl, 2)

		assert.Equal(t, [][]int{{0, 1}, {2, 3}, {4}}, split)
	})

	t.Run("multiple items and zero chunkSize", func(t *testing.T) {
		sl := []int{0, 1, 2, 3, 4}

		split := utils.SplitSliceToChunkIntegers(sl, 0)

		assert.Equal(t, [][]int{{0}, {1}, {2}, {3}, {4}}, split)
	})
}
