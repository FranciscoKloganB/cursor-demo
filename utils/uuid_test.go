//go:build unit

package utils_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"encore.app/utils"
)

func TestParseOrGenerateUUID_GivenValidUUID_ReturnsParsedUUID(t *testing.T) {
	input := "123e4567-e89b-12d3-a456-426614174000"
	expected, _ := uuid.Parse(input)

	actual, err := utils.ParseOrGenerateUUID(input)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestParseOrGenerateUUID_GivenInvalidUUID_ReturnsError(t *testing.T) {
	input := "invalid-uuid"

	_, err := utils.ParseOrGenerateUUID(input)

	assert.Error(t, err)
}

func TestParseOrGenerateUUID_GivenEmptyString_ReturnsError(t *testing.T) {
	input := ""

	_, err := utils.ParseOrGenerateUUID(input)

	assert.Error(t, err)
}
