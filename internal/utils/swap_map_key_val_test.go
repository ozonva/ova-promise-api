package utils_test

import (
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/ozonva/ova-promise-api/internal/utils"
)

func TestSwapMapKeyVal(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		m := make(map[interface{}]interface{})
		res := utils.SwapMapKeyVal(m)
		assert.Equal(t, map[interface{}]interface{}{}, res)

		// check that the original map has not changed
		assert.Equal(t, m, map[interface{}]interface{}{})
	})

	t.Run("one item", func(t *testing.T) {
		m := map[interface{}]interface{}{"one": 1}
		res := utils.SwapMapKeyVal(m)
		assert.Equal(t, map[interface{}]interface{}{1: "one"}, res)

		// check that the original map has not changed
		assert.Equal(t, m, map[interface{}]interface{}{"one": 1})
	})

	t.Run("multiple items", func(t *testing.T) {
		m := map[interface{}]interface{}{"one": 1, 42: "answer"}
		res := utils.SwapMapKeyVal(m)
		assert.Equal(t, map[interface{}]interface{}{1: "one", "answer": 42}, res)

		// check that the original map has not changed
		assert.Equal(t, m, map[interface{}]interface{}{"one": 1, 42: "answer"})
	})
}

//nolint:dupl // tests
func TestSwapMapKeyValIntegers(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		m := make(map[int]int)
		res := utils.SwapMapKeyValIntegers(m)
		assert.Equal(t, map[int]int{}, res)

		// check that the original map has not changed
		assert.Equal(t, m, map[int]int{})
	})

	t.Run("one item", func(t *testing.T) {
		m := map[int]int{1: 10}
		res := utils.SwapMapKeyValIntegers(m)
		assert.Equal(t, map[int]int{10: 1}, res)

		// check that the original map has not changed
		assert.Equal(t, m, map[int]int{1: 10})
	})

	t.Run("empty", func(t *testing.T) {
		m := map[int]int{1: 10, 2: 20}
		res := utils.SwapMapKeyValIntegers(m)
		assert.Equal(t, map[int]int{10: 1, 20: 2}, res)

		// check that the original map has not changed
		assert.Equal(t, m, map[int]int{1: 10, 2: 20})
	})
}

//nolint:dupl // tests
func TestSwapMapKeyValStrings(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		m := make(map[string]string)
		res := utils.SwapMapKeyValStrings(m)
		assert.Equal(t, map[string]string{}, res)

		// check that the original map has not changed
		assert.Equal(t, m, map[string]string{})
	})

	t.Run("one item", func(t *testing.T) {
		m := map[string]string{"1": "10"}
		res := utils.SwapMapKeyValStrings(m)
		assert.Equal(t, map[string]string{"10": "1"}, res)

		// check that the original map has not changed
		assert.Equal(t, m, map[string]string{"1": "10"})
	})

	t.Run("empty", func(t *testing.T) {
		m := map[string]string{"1": "10", "2": "20"}
		res := utils.SwapMapKeyValStrings(m)
		assert.Equal(t, map[string]string{"10": "1", "20": "2"}, res)

		// check that the original map has not changed
		assert.Equal(t, m, map[string]string{"1": "10", "2": "20"})
	})
}
