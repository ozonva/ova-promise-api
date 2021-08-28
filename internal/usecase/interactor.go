package usecase

import (
	"context"

	"go.uber.org/zap"

	"github.com/ozonva/ova-promise-api/internal/domain"
)

type interactor struct {
	promiseRepo PromiseRepository
	logger      *zap.Logger
	chunkSize   int
}

type PromiseRepository interface {
	SavePromise(ctx context.Context, promise *domain.Promise) error
	SavePromiseList(ctx context.Context, promises []domain.Promise) error
	GetPromiseByID(ctx context.Context, id domain.ID) (*domain.Promise, error)
	GetPromiseList(ctx context.Context, limit, offset uint64) ([]domain.Promise, error)
	RemovePromise(ctx context.Context, id domain.ID) error
}
