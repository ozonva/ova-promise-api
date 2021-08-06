package utils_test

import (
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/ozonva/ova-promise-api/internal/utils"
)

func TestFilterSliceByExcluded(t *testing.T) {
	t.Run("empty empty", func(t *testing.T) {
		var (
			initial  []interface{}
			excluded []interface{}
		)

		result := utils.FilterSliceByExcluded(initial, excluded)

		assert.Equal(t, 0, len(result))
	})

	t.Run("one empty", func(t *testing.T) {
		var excluded []interface{}

		initial := []interface{}{"1"}

		result := utils.FilterSliceByExcluded(initial, excluded)

		assert.Equal(t, []interface{}{"1"}, result)
	})

	t.Run("empty one", func(t *testing.T) {
		var initial []interface{}

		excluded := []interface{}{"1"}

		result := utils.FilterSliceByExcluded(initial, excluded)

		assert.Equal(t, 0, len(result))
	})

	t.Run("multiple multiple", func(t *testing.T) {
		excluded := []interface{}{1, "1"}

		initial := []interface{}{1, "1", 3}

		result := utils.FilterSliceByExcluded(initial, excluded)

		assert.Equal(t, []interface{}{3}, result)
	})
}

//nolint:dupl // tests
func TestFilterSliceByExcludedIntegers(t *testing.T) {
	t.Run("empty empty", func(t *testing.T) {
		var (
			initial  []int
			excluded []int
		)

		result := utils.FilterSliceByExcludedIntegers(initial, excluded)

		assert.Equal(t, 0, len(result))
	})

	t.Run("one empty", func(t *testing.T) {
		var excluded []int

		initial := []int{1}

		result := utils.FilterSliceByExcludedIntegers(initial, excluded)

		assert.Equal(t, []int{1}, result)
	})

	t.Run("empty one", func(t *testing.T) {
		var initial []int

		excluded := []int{1}

		result := utils.FilterSliceByExcludedIntegers(initial, excluded)

		assert.Equal(t, 0, len(result))
	})

	t.Run("multiple multiple", func(t *testing.T) {
		excluded := []int{1, 2}

		initial := []int{1, 2, 3}

		result := utils.FilterSliceByExcludedIntegers(initial, excluded)

		assert.Equal(t, []int{3}, result)
	})
}

//nolint:dupl // tests
func TestFilterSliceByExcludedStrings(t *testing.T) {
	t.Run("empty empty", func(t *testing.T) {
		var (
			initial  []string
			excluded []string
		)

		result := utils.FilterSliceByExcludedStrings(initial, excluded)

		assert.Equal(t, 0, len(result))
	})

	t.Run("one empty", func(t *testing.T) {
		var excluded []string

		initial := []string{"1"}

		result := utils.FilterSliceByExcludedStrings(initial, excluded)

		assert.Equal(t, []string{"1"}, result)
	})

	t.Run("empty one", func(t *testing.T) {
		var initial []string

		excluded := []string{"1"}

		result := utils.FilterSliceByExcludedStrings(initial, excluded)

		assert.Equal(t, 0, len(result))
	})

	t.Run("multiple multiple", func(t *testing.T) {
		excluded := []string{"1", "2"}

		initial := []string{"1", "2", "3"}

		result := utils.FilterSliceByExcludedStrings(initial, excluded)

		assert.Equal(t, []string{"3"}, result)
	})
}
