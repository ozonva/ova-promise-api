package domain

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

type UserID = int64

func GenerateID() ID {
	return uuid.New()
}
