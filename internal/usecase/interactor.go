package usecase

import (
	"context"

	"go.uber.org/zap"

	"github.com/ozonva/ova-promise-api/internal/domain"
)

type interactor struct {
	promiseRepo   PromiseRepository
	eventProducer EventProducer
	metrics       ServerMetrics
	logger        *zap.Logger
	chunkSize     int
}

type Transaction interface{}

type Repository interface {
	TransactionCreate(ctx context.Context) (Transaction, error)
	TransactionCommit(ctx context.Context, transaction Transaction) error
	TransactionRollback(ctx context.Context, transaction Transaction) error
}

type PromiseRepository interface {
	Repository
	SavePromise(ctx context.Context, transaction Transaction, promise *domain.Promise) error
	SavePromiseList(ctx context.Context, transaction Transaction, promises []domain.Promise) error
	GetPromiseByID(ctx context.Context, id domain.ID) (*domain.Promise, error)
	GetPromiseList(ctx context.Context, limit, offset uint64) ([]domain.Promise, error)
	RemovePromise(ctx context.Context, transaction Transaction, id domain.ID) error
}

type EventProducer interface {
	SendEventPromiseCreated(ctx context.Context, promise *domain.Promise) error
	SendEventPromiseRemoved(ctx context.Context, id domain.ID) error
	SendEventPromiseUpdated(ctx context.Context, id domain.ID) error
}
