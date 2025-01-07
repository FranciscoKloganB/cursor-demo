//go:build unit

package utils_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"encore.app/utils"
)

func TestFormatNullableTime(t *testing.T) {
	t.Run("returns nil when time is nil", func(t *testing.T) {
		var tNil *time.Time
		result := utils.FormatNullableTime(tNil)
		assert.Nil(t, result, "Expected nil for nil time")
	})

	t.Run("returns formatted time string pointer when time is not nil", func(t *testing.T) {
		timeValue := time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)
		result := utils.FormatNullableTime(&timeValue)
		expected := "2023-10-01T12:00:00Z"
		assert.NotNil(t, result, "Expected non-nil result for valid time")
		assert.Equal(t, &expected, result, "Expected formatted time string pointer")
	})
}
