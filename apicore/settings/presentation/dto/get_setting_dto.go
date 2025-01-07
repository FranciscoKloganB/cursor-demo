package dto

import (
	"time"

	"github.com/google/uuid"
)

// GetSettingResponse represents the response for getting a setting
type GetSettingResponse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Slug      string     `json:"slug"`
	Hint      string     `json:"hint"`
	IsEnabled bool       `json:"is_enabled"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy uuid.UUID  `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *uuid.UUID `json:"updated_by,omitempty"`
	Version   int32      `json:"version"`
}
