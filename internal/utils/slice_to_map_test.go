package utils_test

import (
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/ozonva/ova-promise-api/internal/domain"
	"github.com/ozonva/ova-promise-api/internal/utils"
)

func TestSliceToMapPromises(t *testing.T) {
	p1, _ := domain.NewPromise(domain.GenerateID(), 1, "desc1", nil)
	p2, _ := domain.NewPromise(domain.GenerateID(), 2, "desc2", nil)
	p3, _ := domain.NewPromise(domain.GenerateID(), 3, "desc3", nil)

	sl := []domain.Promise{*p1, *p2, *p3}

	t.Run("empty", func(t *testing.T) {
		res := utils.SliceToMapPromises([]domain.Promise{})

		assert.Equal(t, len(res), 0)
	})

	t.Run("multiple", func(t *testing.T) {
		res := utils.SliceToMapPromises(sl)

		assert.Equal(t, len(res), len(sl))

		for _, p := range sl {
			promise, ok := res[p.ID]

			assert.Equal(t, ok, true)
			assert.Equal(t, p, promise)
		}
	})
}
