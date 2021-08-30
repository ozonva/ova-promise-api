package usecase

import (
	"context"
	"log"

	"go.uber.org/zap"

	"github.com/ozonva/ova-promise-api/internal/domain"
)

const defaultChunkSize = 1

type Handler interface {
	PromiseSave(ctx context.Context, promise *domain.Promise) error
	PromiseSaveList(ctx context.Context, promises []domain.Promise) error
	PromiseGetByID(ctx context.Context, id domain.ID) (*domain.Promise, error)
	PromiseGetList(ctx context.Context, limit, offset uint64) ([]domain.Promise, error)
	PromiseRemove(ctx context.Context, id domain.ID) error
	Flusher
}

type Flusher interface {
	Flush(ctx context.Context, promises []domain.Promise) []domain.Promise
}

type HandlerConstructor struct {
	PromiseRepository PromiseRepository
	ChunkSize         int
	Logger            *zap.Logger
}

func (c HandlerConstructor) New() Handler {
	if c.PromiseRepository == nil {
		log.Fatal("PromiseRepository not set!")
	}

	if c.ChunkSize < 1 {
		log.Println("using default flusher chunk size")

		c.ChunkSize = defaultChunkSize
	}

	return interactor{
		promiseRepo: c.PromiseRepository,
		logger:      c.Logger,
		chunkSize:   c.ChunkSize,
	}
}
