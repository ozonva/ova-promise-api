package usecase_test

import (
	"context"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zaptest"

	"github.com/ozonva/ova-promise-api/internal/domain"
	"github.com/ozonva/ova-promise-api/internal/mocks"
	"github.com/ozonva/ova-promise-api/internal/testdata"
	"github.com/ozonva/ova-promise-api/internal/usecase"
)

func TestInteractor_PromiseSave(t *testing.T) {
	mockPromiseRepo := mocks.PromiseRepository{}
	mockPromiseRepo.On("SavePromise", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("*domain.Promise")).Return(
		func(ctx context.Context, p *domain.Promise) error {
			switch p.ID.String() {
			case testdata.ID0.String():
				return testdata.ErrRepoError
			default:
				return nil
			}
		},
	)

	logger := zaptest.NewLogger(t)
	ucHandler := usecase.HandlerConstructor{
		PromiseRepository: &mockPromiseRepo,
		Logger:            logger,
	}.New()

	t.Run("save with error", func(t *testing.T) {
		err := ucHandler.PromiseSave(context.Background(), &testdata.TestPromiseBroken)

		assert.Equal(t, err, domain.ErrTechnical)
	})

	t.Run("save without error", func(t *testing.T) {
		err := ucHandler.PromiseSave(context.Background(), &testdata.TestPromise1)

		assert.Equal(t, err, nil)
	})
}

func TestInteractor_PromiseSaveList(t *testing.T) {
	mockPromiseRepo := mocks.PromiseRepository{}
	mockPromiseRepo.On("SavePromiseList", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("[]domain.Promise")).Return(
		func(ctx context.Context, promises []domain.Promise) error {
			for _, p := range promises {
				if p.ID.String() == testdata.TestPromiseBroken.ID.String() {
					return testdata.ErrRepoError
				}
			}

			return nil
		},
	)

	logger := zaptest.NewLogger(t)
	ucHandler := usecase.HandlerConstructor{
		PromiseRepository: &mockPromiseRepo,
		Logger:            logger,
	}.New()

	t.Run("save with error", func(t *testing.T) {
		var promises = []domain.Promise{testdata.TestPromiseBroken, testdata.TestPromise1}
		err := ucHandler.PromiseSaveList(context.Background(), promises)

		assert.Equal(t, err, domain.ErrTechnical)
	})

	t.Run("save without error", func(t *testing.T) {
		var promises = []domain.Promise{testdata.TestPromise1, testdata.TestPromise2}
		err := ucHandler.PromiseSaveList(context.Background(), promises)

		assert.Equal(t, err, nil)
	})
}

func TestInteractor_PromiseGetByID(t *testing.T) {
	mockPromiseRepo := mocks.PromiseRepository{}
	mockPromiseRepo.On("GetPromiseByID", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("uuid.UUID")).Return(
		func(ctx context.Context, id domain.ID) *domain.Promise {
			switch id.String() {
			case testdata.ID0.String():
				return nil
			default:
				for _, p := range testdata.PromiseList {
					if id.String() == p.ID.String() {
						return &p
					}
				}

				return nil
			}
		},
		func(ctx context.Context, id domain.ID) error {
			switch id.String() {
			case testdata.ID0.String():
				return testdata.ErrRepoError
			default:
				return nil
			}
		},
	)

	logger := zaptest.NewLogger(t)
	ucHandler := usecase.HandlerConstructor{
		PromiseRepository: &mockPromiseRepo,
		Logger:            logger,
	}.New()

	t.Run("get with error", func(t *testing.T) {
		p, err := ucHandler.PromiseGetByID(context.Background(), testdata.ID0)

		assert.Equal(t, err, domain.ErrTechnical)
		assert.Equal(t, p, nil)
	})

	t.Run("get without error", func(t *testing.T) {
		p, err := ucHandler.PromiseGetByID(context.Background(), testdata.ID1)

		assert.Equal(t, err, nil)
		assert.Equal(t, p, testdata.TestPromise1)
	})
}

func TestInteractor_PromiseGetList(t *testing.T) {
	mockPromiseRepo := mocks.PromiseRepository{}
	mockPromiseRepo.On(
		"GetPromiseList",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("uint64"),
		mock.AnythingOfType("uint64"),
	).Return(
		func(ctx context.Context, limit, offset uint64) []domain.Promise {
			switch limit {
			case 100500:
				return nil
			default:
				return testdata.PromiseList
			}
		},
		func(ctx context.Context, limit, offset uint64) error {
			switch limit {
			case 100500:
				return testdata.ErrRepoError
			default:
				return nil
			}
		},
	)

	logger := zaptest.NewLogger(t)
	ucHandler := usecase.HandlerConstructor{
		PromiseRepository: &mockPromiseRepo,
		Logger:            logger,
	}.New()

	t.Run("get with error", func(t *testing.T) {
		p, err := ucHandler.PromiseGetList(context.Background(), 100500, 0)

		assert.Equal(t, err, domain.ErrTechnical)
		assert.Equal(t, p, nil)
	})

	t.Run("get without error", func(t *testing.T) {
		p, err := ucHandler.PromiseGetList(context.Background(), 3, 0)

		assert.Equal(t, err, nil)
		assert.Equal(t, p, testdata.PromiseList)
	})
}

func TestInteractor_Flush(t *testing.T) {
	mockPromiseRepo := mocks.PromiseRepository{}
	mockPromiseRepo.On("SavePromiseList", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("[]domain.Promise")).Return(
		func(ctx context.Context, promises []domain.Promise) error {
			for _, p := range promises {
				if p.ID.String() == testdata.TestPromiseBroken.ID.String() {
					return testdata.ErrRepoError
				}
			}

			return nil
		},
	)

	logger := zaptest.NewLogger(t)
	ucHandler := usecase.HandlerConstructor{
		PromiseRepository: &mockPromiseRepo,
		Logger:            logger,
	}.New()

	t.Run("flush with error", func(t *testing.T) {
		promises := []domain.Promise{
			testdata.TestPromise3,
			testdata.TestPromiseBroken,
			testdata.TestPromise1,
			testdata.TestPromise2,
		}

		res := ucHandler.Flush(context.Background(), promises)

		assert.Equal(t, res, []domain.Promise{testdata.TestPromiseBroken, testdata.TestPromise1, testdata.TestPromise2})
	})

	t.Run("flush without error", func(t *testing.T) {
		promises := []domain.Promise{
			testdata.TestPromise3,
			testdata.TestPromise1,
			testdata.TestPromise2,
		}

		res := ucHandler.Flush(context.Background(), promises)

		assert.Equal(t, res, nil)
	})
}
