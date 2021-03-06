package utils_test

import (
	"testing"

	"github.com/ozonva/ova-promise-api/internal/domain"

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

func TestSplitToChunk(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		split := utils.SplitSliceToChunk([]interface{}{}, 2)

		assert.Equal(t, nil, split)
	})

	t.Run("one item", func(t *testing.T) {
		split := utils.SplitSliceToChunk([]interface{}{"hello"}, 2)

		assert.Equal(t, [][]interface{}{{"hello"}}, split)
	})

	t.Run("multiple items", func(t *testing.T) {
		sl := []interface{}{"a", "b", "c", 1, 2}

		split := utils.SplitSliceToChunk(sl, 2)

		assert.Equal(t, [][]interface{}{{"a", "b"}, {"c", 1}, {2}}, split)
	})

	t.Run("multiple items and zero chunkSize", func(t *testing.T) {
		sl := []interface{}{"a", "b", "c", 1, 2}

		split := utils.SplitSliceToChunk(sl, 0)

		assert.Equal(t, [][]interface{}{{"a"}, {"b"}, {"c"}, {1}, {2}}, split)
	})
}

func TestSplitSliceToChunkPromises(t *testing.T) {
	p1, _ := domain.NewPromise(domain.GenerateID(), 1, "desc1", nil)
	p2, _ := domain.NewPromise(domain.GenerateID(), 2, "desc2", nil)
	p3, _ := domain.NewPromise(domain.GenerateID(), 3, "desc3", nil)

	t.Run("empty", func(t *testing.T) {
		split := utils.SplitSliceToChunkPromises([]domain.Promise{}, 2)

		assert.Equal(t, nil, split)
	})

	t.Run("one item", func(t *testing.T) {
		split := utils.SplitSliceToChunkPromises([]domain.Promise{*p1}, 2)

		assert.Equal(t, [][]domain.Promise{{*p1}}, split)
	})

	t.Run("multiple items", func(t *testing.T) {
		sl := []domain.Promise{*p1, *p2, *p3}

		split := utils.SplitSliceToChunkPromises(sl, 2)

		assert.Equal(t, [][]domain.Promise{{*p1, *p2}, {*p3}}, split)
	})

	t.Run("multiple items and zero chunkSize", func(t *testing.T) {
		sl := []domain.Promise{*p1, *p2, *p3}

		split := utils.SplitSliceToChunkPromises(sl, 0)

		assert.Equal(t, [][]domain.Promise{{*p1}, {*p2}, {*p3}}, split)
	})
}
