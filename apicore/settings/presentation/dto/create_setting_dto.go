package dto

import (
	"time"

	"encore.app/apicore/common/pipes/request"
	"github.com/google/uuid"
)

// CreateSettingRequest represents the payload for creating a new setting
type CreateSettingRequest struct {
	// The name to be given to the setting
	Name string `json:"name" mod:"trim,title" validate:"required,min=3,max=100"`
	// Unique identifier for code references, typically in kebab-case. When not provided uses name property converted to kebab-case.
	Slug string `json:"slug" mod:"trim,slug" validate:"omitempty,min=3,max=50"`
	// Concise explanation of the setting's purpose and behavior
	Hint string `json:"hint" mod:"trim,ucfirst" validate:"required,min=10,max=500"`
	// Determines if the setting should be active upon creation
	IsEnabled bool `json:"is_enabled"`
}

// Validate implements the validation for GetSettingsRequest
func (r *CreateSettingRequest) Validate() error {
	return request.ParseRequest(r)
}

// CreateSettingResponse represents the response after creating a setting
type CreateSettingResponse struct {
	// Unique system-generated identifier
	ID string `json:"id"`
	// Display identifier shown in the UI
	Name string `json:"name"`
	// Unique identifier for code references
	Slug string `json:"slug"`
	// Concise explanation of the setting's purpose and behavior
	Hint string `json:"hint"`
	// Current activation status
	IsEnabled bool `json:"is_enabled"`
	// Timestamp of setting creation
	CreatedAt time.Time `json:"created_at"`
	// ID of the actor who created the setting creation
	CreatedBy uuid.UUID `json:"created_by"`
	// Timestamp of the last update
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// ID of the actor who last updated or created the setting
	UpdatedBy *uuid.UUID `json:"updated_by,omitempty"`
	// The object version number
	Version int32 `json:"version"`
}
