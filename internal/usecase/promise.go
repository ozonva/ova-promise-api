package usecase

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/ozonva/ova-promise-api/internal/domain"
	"github.com/ozonva/ova-promise-api/internal/utils"
)

func (i interactor) PromiseSave(ctx context.Context, promise *domain.Promise) error {
	if err := i.promiseRepo.SavePromise(ctx, promise); err != nil {
		i.logger.Error(
			"usecase promise-save: error while saving object to repo",
			zap.Any("promise", promise),
			zap.Error(err),
		)

		return domain.ErrTechnical
	}

	return nil
}

func (i interactor) PromiseSaveList(ctx context.Context, promises []domain.Promise) error {
	if err := i.promiseRepo.SavePromiseList(ctx, promises); err != nil {
		i.logger.Error(
			"usecase promise-save-list: error while saving list of object to repo",
			zap.Any("promises", promises),
			zap.Error(err),
		)

		return domain.ErrTechnical
	}

	return nil
}

func (i interactor) PromiseGetByID(ctx context.Context, id domain.ID) (*domain.Promise, error) {
	promise, err := i.promiseRepo.GetPromiseByID(ctx, id)

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, err
		}

		i.logger.Error(
			"usecase promise-get-by-id: error while reading object from repo",
			zap.String("id", id.String()),
			zap.Error(err),
		)

		return nil, domain.ErrTechnical
	}

	return promise, nil
}

func (i interactor) PromiseGetList(ctx context.Context, limit, offset uint64) ([]domain.Promise, error) {
	res, err := i.promiseRepo.GetPromiseList(ctx, limit, offset)

	if err != nil {
		i.logger.Error(
			"usecase promise-get-list: error while reading list of object from repo",
			zap.Uint64("limit", limit),
			zap.Uint64("offset", offset),
			zap.Error(err),
		)

		return nil, domain.ErrTechnical
	}

	return res, nil
}

func (i interactor) Flush(ctx context.Context, promises []domain.Promise) []domain.Promise {
	var result []domain.Promise

	chunks := utils.SplitSliceToChunkPromises(promises, i.chunkSize)

	for idx, chunk := range chunks {
		if err := i.promiseRepo.SavePromiseList(ctx, chunk); err != nil {
			i.logger.Error(
				"usecase flush: error while flushing objects to repo",
				zap.Any("promises-chunk", chunk),
				zap.Error(err),
			)

			for j := idx; j < len(chunks); j++ {
				result = append(result, chunks[j]...)
			}

			return result
		}
	}

	return result
}
