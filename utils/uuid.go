package utils

import (
	"errors"

	"github.com/google/uuid"
)

// ParseOrGenerateUUID attempts to parse the given string as a UUID.
// If the string is empty, it returns an error.
func ParseOrGenerateUUID(id string) (uuid.UUID, error) {
	if id == "" {
		return uuid.UUID{}, errors.New("empty string is not a valid UUID")
	}

	parsedUUID, err := uuid.Parse(id)

	if err != nil {
		return uuid.UUID{}, errors.New("invalid UUID: " + err.Error())
	}

	return parsedUUID, nil
}
