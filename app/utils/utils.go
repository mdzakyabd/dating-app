package utils

import (
	"github.com/google/uuid"
)

// DistinctUUIDs returns a new slice with only distinct UUIDs from the input slice.
func DistinctUUIDs(input []uuid.UUID) []uuid.UUID {
	uniqueMap := make(map[uuid.UUID]bool)
	uniqueList := []uuid.UUID{}

	for _, id := range input {
		if _, exists := uniqueMap[id]; !exists {
			uniqueMap[id] = true
			uniqueList = append(uniqueList, id)
		}
	}

	return uniqueList
}
