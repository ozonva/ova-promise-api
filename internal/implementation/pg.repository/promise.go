package pgrepo

import (
	"context"
	"errors"

	"github.com/ozonva/ova-promise-api/internal/usecase"

	"github.com/jackc/pgx/v4"

	"github.com/ozonva/ova-promise-api/internal/domain"
)

func (r rw) SavePromise(ctx context.Context, transaction usecase.Transaction, p *domain.Promise) error {
	tx, ok := transaction.(pgx.Tx)
	if !ok {
		return ErrInvalidTransaction
	}

	if _, err := tx.Exec(
		ctx,
		`insert into promises (id, user_id, description, status, date_deadline, created_at, updated_at)
			values ($1, $2, $3, $4, $5, $6, $7)
			on conflict (id)
			do update
			set description=$3, status=$4, date_deadline=$5, updated_at=$7`,
		p.ID, p.UserID, p.Description, p.Status, p.DateDeadline, p.CreatedAt, p.UpdatedAt,
	); err != nil {
		return err
	}

	return nil
}

func (r rw) SavePromiseList(ctx context.Context, transaction usecase.Transaction, promises []domain.Promise) error {
	tx, ok := transaction.(pgx.Tx)
	if !ok {
		return ErrInvalidTransaction
	}

	batch := &pgx.Batch{}

	for _, p := range promises {
		batch.Queue(
			`insert into promises (id, user_id, description, status, date_deadline, created_at, updated_at) 
			values ($1, $2, $3, $4, $5, $6, $7)`,
			p.ID, p.UserID, p.Description, p.Status, p.DateDeadline, p.CreatedAt, p.UpdatedAt,
		)
	}

	br := tx.SendBatch(ctx, batch)

	if err := br.Close(); err != nil {
		return err
	}

	return nil
}

func (r rw) GetPromiseByID(ctx context.Context, id domain.ID) (*domain.Promise, error) {
	p := domain.Promise{}

	err := r.store.QueryRow(
		ctx,
		`select id, user_id, description, status, date_deadline, created_at, updated_at
		from promises where id=$1`, id,
	).Scan(
		&p.ID, &p.UserID, &p.Description, &p.Status, &p.DateDeadline, &p.CreatedAt, &p.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}

		return nil, err
	}

	return &p, nil
}

func (r rw) GetPromiseList(ctx context.Context, limit, offset uint64) ([]domain.Promise, error) {
	result := make([]domain.Promise, 0)

	rows, _ := r.store.Query(
		ctx,
		`select id, user_id, description, status, date_deadline, created_at, updated_at from promises limit $1 offset $2`,
		limit, offset,
	)

	for rows.Next() {
		p := domain.Promise{}

		if err := rows.Scan(&p.ID, &p.UserID, &p.Description, &p.Status, &p.DateDeadline, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}

		result = append(result, p)
	}

	return result, nil
}

func (r rw) RemovePromise(ctx context.Context, transaction usecase.Transaction, id domain.ID) error {
	tx, ok := transaction.(pgx.Tx)
	if !ok {
		return ErrInvalidTransaction
	}

	if _, err := tx.Exec(ctx, `delete from promises where id=$1`, id); err != nil {
		return err
	}

	return nil
}
