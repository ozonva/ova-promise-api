package usecase

import (
	"context"

	"go.uber.org/zap"

	"github.com/ozonva/ova-promise-api/internal/domain"
)

type interactor struct {
	promiseRepo   PromiseRepository
	eventProducer EventProducer
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
	NewEventPromiseCreated(ctx context.Context, promise *domain.Promise) error
	NewEventPromiseRemoved(ctx context.Context, id domain.ID) error
	NewEventPromiseUpdated(ctx context.Context, id domain.ID) error
}
