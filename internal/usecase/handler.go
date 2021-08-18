package usecase

import (
	"context"
	"log"

	"go.uber.org/zap"

	"github.com/ozonva/ova-promise-api/internal/domain"
)

type Handler interface {
	PromiseSave(ctx context.Context, promise *domain.Promise) error
	PromiseSaveList(ctx context.Context, promises []domain.Promise) error
	PromiseGetByID(ctx context.Context, id domain.ID) (*domain.Promise, error)
	PromiseGetList(ctx context.Context, limit, offset uint64) ([]domain.Promise, error)
	Flusher
}

type Flusher interface {
	Flush(ctx context.Context, promises []domain.Promise) []domain.Promise
}

type HandlerConstructor struct {
	PromiseRepository PromiseRepository
	Logger            *zap.Logger
}

func (c HandlerConstructor) New(chunkSize int) Handler {
	if c.PromiseRepository == nil {
		log.Fatal("PromiseRepository not set!")
	}

	return interactor{
		promiseRepo: c.PromiseRepository,
		logger:      c.Logger,
		chunkSize:   chunkSize,
	}
}
