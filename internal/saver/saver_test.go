package saver_test

import (
	"context"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"

	"github.com/ozonva/ova-promise-api/internal/domain"
	"github.com/ozonva/ova-promise-api/internal/mocks"
	"github.com/ozonva/ova-promise-api/internal/saver"
	"github.com/ozonva/ova-promise-api/internal/testdata"
)

func TestSaver_Save(t *testing.T) {
	ucHandler := mocks.Handler{}
	ucHandler.On("Flush", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("[]domain.Promise")).Return(
		func(ctx context.Context, promises []domain.Promise) []domain.Promise {
			for _, p := range promises {
				if p.ID.String() == testdata.TestPromiseBroken.ID.String() {
					return promises
				}
			}

			return nil
		},
	)

	t.Run("not enough capacity", func(t *testing.T) {
		s := saver.NewSaver(context.Background(), 1, 2, &ucHandler)
		err := s.Save(testdata.TestPromise1)
		assert.Equal(t, err, nil)
		err = s.Save(testdata.TestPromise2)
		assert.Equal(t, err, nil)
		err = s.Save(testdata.TestPromise3)
		assert.Equal(t, err, saver.ErrFullBuffer)
	})
}
