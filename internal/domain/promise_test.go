package domain_test

import (
	"testing"
	"time"

	"github.com/go-playground/assert/v2"

	"github.com/ozonva/ova-promise-api/internal/domain"
)

var (
	promiseID   = domain.GenerateID()
	userID      = domain.UserID(1)
	description = "description"
	nowTime     = time.Now().UTC()
	tomorrow    = time.Now().AddDate(0, 0, 1).UTC()
	yesterday   = time.Now().AddDate(0, 0, -1).UTC()

	promise = &domain.Promise{
		ID:          promiseID,
		UserID:      userID,
		Description: description,
		Status:      domain.PromiseStatusValueDraft.String(),
		CreatedAt:   nowTime,
		UpdatedAt:   nowTime,
	}
)

func TestNewPromise(t *testing.T) {
	t.Run("normally create", func(t *testing.T) {
		p, err := domain.NewPromise(promiseID, userID, description, nil)
		assert.Equal(t, nil, err)
		assert.Equal(t, promiseID, p.ID)
		assert.Equal(t, description, p.Description)
		assert.Equal(t, domain.PromiseStatusValueDraft.String(), p.Status)
		assert.Equal(t, nil, p.DateDeadline)
		assert.Equal(t, p.CreatedAt, p.UpdatedAt)
		assert.Equal(t, true, time.Now().UTC().After(p.CreatedAt))
		assert.Equal(t, true, time.Now().UTC().After(p.UpdatedAt))

		p, err = domain.NewPromise(promiseID, userID, description, &tomorrow)
		assert.Equal(t, nil, err)
		assert.Equal(t, p.DateDeadline, tomorrow)
	})

	t.Run("yesterday date", func(t *testing.T) {
		p, err := domain.NewPromise(promiseID, userID, description, &yesterday)
		assert.Equal(t, err, domain.ErrEarlierDateDeadline)
		assert.Equal(t, nil, p)
	})

	t.Run("empty description", func(t *testing.T) {
		p, err := domain.NewPromise(promiseID, userID, "", nil)
		assert.Equal(t, err, domain.ErrEmptyDescription)
		assert.Equal(t, nil, p)
	})
}

func TestUpdatePromise(t *testing.T) {
	t.Run("update invalid status", func(t *testing.T) {
		p := promise
		err := domain.UpdatePromise(p, domain.SetPromiseStatus("bad_status"))
		assert.Equal(t, err, domain.ErrInvalidPromiseStatus)
		assert.Equal(t, p.CreatedAt, p.UpdatedAt)
		assert.Equal(t, p.UpdatedAt, nowTime)
	})

	t.Run("update status", func(t *testing.T) {
		p := promise
		err := domain.UpdatePromise(p, domain.SetPromiseStatus("in_progress"))
		assert.Equal(t, err, nil)
		assert.Equal(t, true, p.UpdatedAt.After(p.CreatedAt))
	})

	t.Run("update date deadline invalid", func(t *testing.T) {
		p, _ := domain.NewPromise(promiseID, userID, description, nil)

		err := domain.UpdatePromise(p, domain.SetPromiseDateDeadline(123))
		assert.Equal(t, err, domain.ErrInvalidDateDeadline)
		assert.Equal(t, p.CreatedAt, p.UpdatedAt)

		err = domain.UpdatePromise(p, domain.SetPromiseDateDeadline(&yesterday))
		assert.Equal(t, err, domain.ErrEarlierDateDeadline)
		assert.Equal(t, p.CreatedAt, p.UpdatedAt)
	})

	t.Run("update date deadline", func(t *testing.T) {
		p, _ := domain.NewPromise(promiseID, userID, description, nil)

		err := domain.UpdatePromise(p, domain.SetPromiseDateDeadline(&tomorrow))
		assert.Equal(t, nil, err)
		assert.NotEqual(t, nil, p.DateDeadline)
		assert.Equal(t, true, p.UpdatedAt.After(p.CreatedAt))

		var x *time.Time
		err = domain.UpdatePromise(p, domain.SetPromiseDateDeadline(x))
		assert.Equal(t, nil, err)
		assert.Equal(t, nil, p.DateDeadline)
		assert.Equal(t, true, p.UpdatedAt.After(p.CreatedAt))
	})

	t.Run("update description invalid", func(t *testing.T) {
		p, _ := domain.NewPromise(promiseID, userID, description, nil)

		err := domain.UpdatePromise(p, domain.SetPromiseDescription(123))
		assert.Equal(t, err, domain.ErrInvalidDescription)
		assert.Equal(t, p.CreatedAt, p.UpdatedAt)

		err = domain.UpdatePromise(p, domain.SetPromiseDescription(""))
		assert.Equal(t, err, domain.ErrEmptyDescription)
		assert.Equal(t, p.CreatedAt, p.UpdatedAt)
	})
}

func TestCheckIsValidPromiseStatus(t *testing.T) {
	t.Run("invalid promise status", func(t *testing.T) {
		check := domain.CheckIsValidPromiseStatus("invalid_status")
		assert.Equal(t, false, check)
	})

	for _, s := range []string{"draft", "in_progress", "completed"} {
		status := s

		t.Run("promise status", func(t *testing.T) {
			check := domain.CheckIsValidPromiseStatus(status)
			assert.Equal(t, true, check)
		})
	}
}

func TestPromise_Validate(t *testing.T) {
	p := domain.Promise{
		ID:           promiseID,
		UserID:       userID,
		Description:  description,
		Status:       domain.PromiseStatusValueDraft.String(),
		DateDeadline: nil,
		CreatedAt:    nowTime,
		UpdatedAt:    nowTime,
	}

	t.Run("validate invalid description", func(t *testing.T) {
		p.Description = ""
		err := p.Validate()
		assert.Equal(t, err, domain.ErrEmptyDescription)
	})

	t.Run("validate invalid description", func(t *testing.T) {
		p.Description = domain.PromiseStatusValueDraft.String()
		p.Status = ""
		err := p.Validate()
		assert.Equal(t, err, domain.ErrInvalidPromiseStatus)
	})
}
