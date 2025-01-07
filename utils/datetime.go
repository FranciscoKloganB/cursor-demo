package utils

import "time"

// FormatNullableTime formats a nullable time pointer to RFC3339 string or empty string if nil
func FormatNullableTime(t *time.Time) *string {
	if t == nil {
		return nil
	}
	formatted := t.Format(time.RFC3339)
	return &formatted
}
