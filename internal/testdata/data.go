package testdata

import (
	"errors"
	"github.com/google/uuid"
	"github.com/ozonva/ova-promise-api/internal/domain"
	"time"
)

var ErrRepoError = errors.New("repo error")

var ID0, _ = uuid.Parse("00000000-0000-0000-0000-000000000000")
var ID1, _ = uuid.Parse("00000000-0000-0000-0000-000000000001")
var ID2, _ = uuid.Parse("00000000-0000-0000-0000-000000000002")
var ID3, _ = uuid.Parse("00000000-0000-0000-0000-000000000003")

var TestPromise1 = domain.Promise{
	ID:          ID1,
	UserID:      1,
	Description: "Test Promise 1",
	Status:      "draft",
	CreatedAt:   time.Now().UTC(),
	UpdatedAt:   time.Now().UTC(),
}

var TestPromise2 = domain.Promise{
	ID:          ID2,
	UserID:      2,
	Description: "Test Promise 2",
	Status:      "in_progress",
	CreatedAt:   time.Now().UTC(),
	UpdatedAt:   time.Now().UTC(),
}

var TestPromise3 = domain.Promise{
	ID:          ID3,
	UserID:      3,
	Description: "Test Promise 3",
	Status:      "completed",
	CreatedAt:   time.Now().UTC(),
	UpdatedAt:   time.Now().UTC(),
}

var TestPromiseBroken = domain.Promise{
	ID:          ID0,
	UserID:      0,
	Description: "Broken",
	Status:      "draft",
	CreatedAt:   time.Now().UTC(),
	UpdatedAt:   time.Now().UTC(),
}

var PromiseList = []domain.Promise{
	TestPromise1,
	TestPromise2,
	TestPromise3,
}
