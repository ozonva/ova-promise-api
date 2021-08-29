package pgrepo

import (
	"context"

	"github.com/ozonva/ova-promise-api/internal/domain"
)

func (r rw) SavePromise(ctx context.Context, promise *domain.Promise) error {
	panic("implement me")
}

func (r rw) SavePromiseList(ctx context.Context, promises []domain.Promise) error {
	panic("implement me")
}

func (r rw) GetPromiseByID(ctx context.Context, id domain.ID) (*domain.Promise, error) {
	panic("implement me")
}

func (r rw) GetPromiseList(ctx context.Context, limit, offset uint64) ([]domain.Promise, error) {
	panic("implement me")
}
