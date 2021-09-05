package domain

import (
	"time"
)

type Promise struct {
	ID           ID
	UserID       UserID
	Description  string
	Status       string
	DateDeadline *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type PromiseUpdateProperty int

const (
	PromiseDescription PromiseUpdateProperty = iota
	PromiseStatus
	PromiseDateDeadline
)

type PromiseStatusValue struct {
	value string
}

func (p *PromiseStatusValue) String() string {
	return p.value
}

var (
	PromiseStatusValueDraft      = PromiseStatusValue{"draft"}
	PromiseStatusValueInProgress = PromiseStatusValue{"in_progress"}
	PromiseStatusValueCompleted  = PromiseStatusValue{"completed"}
)

var promiseStatusValues = []PromiseStatusValue{
	PromiseStatusValueDraft,
	PromiseStatusValueInProgress,
	PromiseStatusValueCompleted,
}

func CheckIsValidPromiseStatus(s string) bool {
	for _, x := range promiseStatusValues {
		if x.String() == s {
			return true
		}
	}

	return false
}

func (p *Promise) Validate() error {
	if !CheckIsValidPromiseStatus(p.Status) {
		return ErrInvalidPromiseStatus
	}

	if p.Description == "" {
		return ErrEmptyDescription
	}

	return nil
}

// NewPromise return new promise in draft status.
func NewPromise(id ID, userID UserID, description string, deadline *time.Time) (*Promise, error) {
	now := time.Now().UTC()

	p := &Promise{
		ID:        id,
		UserID:    userID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := UpdatePromise(
		p,
		SetPromiseStatus(PromiseStatusValueDraft.String()),
		SetPromiseDescription(description),
		SetPromiseDateDeadline(deadline),
	)
	if err != nil {
		return nil, err
	}

	p.UpdatedAt = p.CreatedAt

	return p, nil
}

func UpdatePromise(initial *Promise, opts ...func(fields *Promise) error) error {
	for _, v := range opts {
		if err := v(initial); err != nil {
			return err
		}
	}

	initial.UpdatedAt = time.Now().UTC()

	return nil
}

func SetPromiseDescription(input interface{}) func(fields *Promise) error {
	return func(initial *Promise) error {
		value, ok := input.(string)

		if !ok {
			return ErrInvalidDescription
		}

		if value == "" {
			return ErrEmptyDescription
		}

		initial.Description = value

		return nil
	}
}

func SetPromiseStatus(input interface{}) func(fields *Promise) error {
	return func(initial *Promise) error {
		value, ok := input.(string)

		if !ok || !CheckIsValidPromiseStatus(value) {
			return ErrInvalidPromiseStatus
		}

		initial.Status = value

		return nil
	}
}

func SetPromiseDateDeadline(input interface{}) func(fields *Promise) error {
	return func(initial *Promise) error {
		value, ok := input.(*time.Time)

		if !ok {
			return ErrInvalidDateDeadline
		}

		if value == nil {
			initial.DateDeadline = value

			return nil
		}

		if value.Before(time.Now().UTC()) {
			return ErrEarlierDateDeadline
		}

		initial.DateDeadline = value

		return nil
	}
}
