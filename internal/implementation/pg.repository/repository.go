package pgrepo

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/ozonva/ova-promise-api/internal/usecase"
)

type rw struct {
	store *pgxpool.Pool
}

func (r rw) TransactionCreate(ctx context.Context) (usecase.Transaction, error) {
	tx, err := r.store.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (r rw) TransactionCommit(ctx context.Context, transaction usecase.Transaction) error {
	tx, ok := transaction.(pgx.Tx)
	if !ok {
		return ErrInvalidTransaction
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (r rw) TransactionRollback(ctx context.Context, transaction usecase.Transaction) error {
	tx, ok := transaction.(pgx.Tx)
	if !ok {
		return ErrInvalidTransaction
	}

	if err := tx.Rollback(ctx); err != nil {
		return err
	}

	return nil
}

func CreateRepository(dbpool *pgxpool.Pool) usecase.PromiseRepository {
	return rw{
		store: dbpool,
	}
}
