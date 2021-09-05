package usecase

import (
	"context"
	"errors"

	"github.com/opentracing/opentracing-go"

	"go.uber.org/zap"

	"github.com/ozonva/ova-promise-api/internal/domain"
	"github.com/ozonva/ova-promise-api/internal/utils"
)

func (i interactor) PromiseSave(ctx context.Context, promise *domain.Promise) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "uc:promise_save")
	defer span.Finish()

	tx, err := i.promiseRepo.TransactionCreate(ctx)
	if err != nil {
		i.logger.Error("usecase promise-save: error while creating repo transaction", zap.Error(err))

		return domain.ErrTechnical
	}

	if err := i.promiseRepo.SavePromise(ctx, tx, promise); err != nil {
		i.logger.Warn("usecase promise-save: transaction rollback")

		if errRollback := i.promiseRepo.TransactionRollback(ctx, tx); errRollback != nil {
			i.logger.Error("usecase promise-save: error while transaction rollback", zap.Error(errRollback))

			return domain.ErrTechnical
		}

		i.logger.Error(
			"usecase promise-save: error while saving object to repo",
			zap.Any("promise", promise),
			zap.Error(err),
		)

		return domain.ErrTechnical
	}

	if err := i.eventProducer.NewEventPromiseCreated(ctx, promise); err != nil {
		i.logger.Error(
			"usecase promise-save: error while sending NewEventPromiseCreated to message broker",
			zap.Any("promise", promise),
			zap.Error(err),
		)

		return domain.ErrTechnical
	}

	if err := i.promiseRepo.TransactionCommit(ctx, tx); err != nil {
		i.logger.Error(
			"usecase promise-save: error while committing transaction",
			zap.Error(err),
		)

		return domain.ErrTechnical
	}

	return nil
}

func (i interactor) PromiseSaveListChunks(ctx context.Context, promises []domain.Promise, chunkSize int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "uc:promise_save_list_chunks")
	defer span.Finish()

	logger := i.logger.With(zap.String("usecase", "PromiseSaveListChunks"))

	if chunkSize == 0 {
		chunkSize = i.chunkSize
	}

	chunks := utils.SplitSliceToChunkPromises(promises, chunkSize)

	for idx, chunk := range chunks {
		if err := i.PromiseSaveList(ctx, chunk); err != nil {
			logger.Error("error while saving chunk", zap.Int("chunk-num", idx))

			return domain.ErrTechnical
		}
	}

	return nil
}

func (i interactor) PromiseSaveList(ctx context.Context, promises []domain.Promise) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "uc:promise_save_list")
	defer span.Finish()

	logger := i.logger.With(zap.String("usecase", "PromiseSaveList"))

	tx, err := i.promiseRepo.TransactionCreate(ctx)
	if err != nil {
		logger.Error("error while creating repo transaction", zap.Error(err))

		return domain.ErrTechnical
	}

	if err := i.promiseRepo.SavePromiseList(ctx, &tx, promises); err != nil {
		logger.Warn("transaction rollback")

		if errRollback := i.promiseRepo.TransactionRollback(ctx, &tx); errRollback != nil {
			logger.Error("error while transaction rollback", zap.Error(errRollback))

			return domain.ErrTechnical
		}

		logger.Error("error while saving list of object to repo", zap.Any("promises", promises), zap.Error(err))

		return domain.ErrTechnical
	}

	if err := i.promiseRepo.TransactionCommit(ctx, &tx); err != nil {
		i.logger.Error("error while committing transaction", zap.Error(err))

		return domain.ErrTechnical
	}

	return nil
}

func (i interactor) PromiseGetByID(ctx context.Context, id domain.ID) (*domain.Promise, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "uc:promise_get_by_id")
	defer span.Finish()

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
	span, ctx := opentracing.StartSpanFromContext(ctx, "uc:promise_get_list")
	defer span.Finish()

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
	span, ctx := opentracing.StartSpanFromContext(ctx, "uc:promise_flush")
	defer span.Finish()

	var result []domain.Promise

	chunks := utils.SplitSliceToChunkPromises(promises, i.chunkSize)

	for idx, chunk := range chunks {
		// each chunk has its own transaction
		tx, err := i.promiseRepo.TransactionCreate(ctx)
		if err != nil {
			i.logger.Error("usecase Flush: error while creating repo transaction", zap.Error(err))

			// current chunk + remaining chunks
			for j := idx; j < len(chunks); j++ {
				result = append(result, chunks[j]...)
			}

			return result
		}

		if err := i.promiseRepo.SavePromiseList(ctx, &tx, chunk); err != nil {
			i.logger.Error(
				"usecase flush: error while flushing objects to repo",
				zap.Any("promises-chunk", chunk),
				zap.Error(err),
			)

			// current chunk + remaining chunks
			for j := idx; j < len(chunks); j++ {
				result = append(result, chunks[j]...)
			}

			i.logger.Warn("usecase flush: transaction rollback")

			if errRollback := i.promiseRepo.TransactionRollback(ctx, &tx); errRollback != nil {
				i.logger.Error("usecase flush: error while transaction rollback", zap.Error(errRollback))
			}

			return result
		}

		if err := i.promiseRepo.TransactionCommit(ctx, &tx); err != nil {
			i.logger.Error("usecase flush: error while committing transaction", zap.Error(err))

			// current chunk + remaining chunks
			for j := idx; j < len(chunks); j++ {
				result = append(result, chunks[j]...)
			}

			return result
		}
	}

	return nil
}

func (i interactor) PromiseUpdate(
	ctx context.Context, id domain.ID, fieldsToUpdate map[domain.PromiseUpdateProperty]interface{},
) (*domain.Promise, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "uc:promise_update")
	defer span.Finish()

	logger := i.logger.With(zap.String("usecase", "PromiseUpdate"), zap.String("promise-id", id.String()))

	promise, err := i.promiseRepo.GetPromiseByID(ctx, id)
	if err != nil {
		if err != nil {
			logger.Error("error while reading promise from repo", zap.Error(err))

			return nil, domain.ErrTechnical
		}
	}

	if err := domain.UpdatePromise(promise,
		domain.SetPromiseDescription(fieldsToUpdate[domain.PromiseDescription]),
		domain.SetPromiseStatus(fieldsToUpdate[domain.PromiseStatus]),
		domain.SetPromiseDateDeadline(fieldsToUpdate[domain.PromiseDateDeadline]),
	); err != nil {
		logger.Error("error while updating promise", zap.Error(err))

		return nil, err
	}

	tx, err := i.promiseRepo.TransactionCreate(ctx)
	if err != nil {
		logger.Error("error while creating repo transaction", zap.Error(err))

		return nil, domain.ErrTechnical
	}

	if err := i.promiseRepo.SavePromise(ctx, tx, promise); err != nil {
		logger.Error("error while saving promise in repo", zap.Error(err))

		return nil, domain.ErrTechnical
	}

	if err := i.eventProducer.NewEventPromiseUpdated(ctx, id); err != nil {
		i.logger.Error(
			"error while sending NewEventPromiseUpdated to message broker",
			zap.Any("promise-id", id),
			zap.Error(err),
		)

		return nil, domain.ErrTechnical
	}

	if err := i.promiseRepo.TransactionCommit(ctx, &tx); err != nil {
		i.logger.Error(
			"error while committing transaction",
			zap.Error(err),
		)

		return nil, domain.ErrTechnical
	}

	return promise, nil
}

func (i interactor) PromiseRemove(ctx context.Context, id domain.ID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "uc:promise_remove")
	defer span.Finish()

	tx, err := i.promiseRepo.TransactionCreate(ctx)
	if err != nil {
		i.logger.Error("usecase promise-remove: error while creating repo transaction", zap.Error(err))

		return domain.ErrTechnical
	}

	if err := i.promiseRepo.RemovePromise(ctx, &tx, id); err != nil {
		i.logger.Error(
			"usecase promise-remove: error while flushing objects to repo",
			zap.String("promises-id", id.String()),
			zap.Error(err),
		)

		i.logger.Warn("usecase promise-remove: transaction rollback")

		if errRollback := i.promiseRepo.TransactionRollback(ctx, &tx); errRollback != nil {
			i.logger.Error("usecase promise-remove: error while transaction rollback", zap.Error(errRollback))

			return domain.ErrTechnical
		}

		return err
	}

	if err := i.eventProducer.NewEventPromiseRemoved(ctx, id); err != nil {
		i.logger.Error(
			"usecase promise-save: error while sending NewEventPromiseRemoved to message broker",
			zap.Any("promise-id", id),
			zap.Error(err),
		)

		return domain.ErrTechnical
	}

	if err := i.promiseRepo.TransactionCommit(ctx, &tx); err != nil {
		i.logger.Error(
			"usecase promise-remove: error while committing transaction",
			zap.Error(err),
		)

		return domain.ErrTechnical
	}

	return nil
}
