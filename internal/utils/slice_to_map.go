package utils

import "github.com/ozonva/ova-promise-api/internal/domain"

// SliceToMapPromises convert slice of domain.Promise to map with key = promise.id and value = promise.
func SliceToMapPromises(initial []domain.Promise) map[domain.ID]domain.Promise {
	result := make(map[domain.ID]domain.Promise)

	for _, p := range initial {
		result[p.ID] = p
	}

	return result
}
