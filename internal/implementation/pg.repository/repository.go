package pgrepo

import (
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/ozonva/ova-promise-api/internal/usecase"
)

type rw struct {
	store *pgxpool.Pool
}

func CreateRepository(dbpool *pgxpool.Pool) usecase.PromiseRepository {
	return rw{
		store: dbpool,
	}
}
