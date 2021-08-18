package domain

import "errors"

var (
	ErrInvalidDescription   = errors.New("invalid description")
	ErrInvalidDateDeadline  = errors.New("invalid date deadline")
	ErrInvalidPromiseStatus = errors.New("invalid promise status")
	ErrEmptyDescription     = errors.New("description can't e empty")
	ErrEarlierDateDeadline  = errors.New("date deadline can't be earlier than now")
	ErrTechnical            = errors.New("some shit happened")
	ErrNotFound             = errors.New("entity not found")
)
