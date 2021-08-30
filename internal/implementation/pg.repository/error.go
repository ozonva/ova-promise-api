package pgrepo

import "errors"

var (
	ErrInvalidTransaction = errors.New("it's not a pgx.Tx")
)
