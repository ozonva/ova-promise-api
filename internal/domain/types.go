package domain

import (
	"time"

	"github.com/google/uuid"
)

type ID = uuid.UUID

type UserID = int64

func GenerateID() ID {
	return uuid.New()
}

type NullTime struct {
	value *time.Time
}

func (t *NullTime) String() string {
	if t != nil {
		if t.value != nil {
			return t.String()
		}

		return ""
	}

	return ""
}

func (t *NullTime) ToTime() *time.Time {
	if t.value == nil {
		return nil
	}

	return t.value
}
