package domain_test

import (
	"testing"

	"github.com/google/uuid"

	"github.com/go-playground/assert/v2"
	"github.com/ozonva/ova-promise-api/internal/domain"
)

func TestGenerateID(t *testing.T) {
	t.Run("generate id", func(t *testing.T) {
		id := domain.GenerateID()
		assert.NotEqual(t, id, nil)
		_, err := uuid.Parse(id.String())
		assert.Equal(t, err, nil)
	})
}
