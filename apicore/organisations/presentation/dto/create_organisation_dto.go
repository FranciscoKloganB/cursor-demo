package dto

import (
	"time"

	"encore.app/apicore/common/pipes/request"
	"github.com/google/uuid"
)

// CreateOrganisationRequest represents the request payload for organization creation.
type CreateOrganisationRequest struct {
	// The name to be assigned to organisation
	Name string `json:"name" mod:"trim,title" validate:"required,min=2,max=128"`
}

// Validate implements the validation for CreateOrganisationRequest
func (r *CreateOrganisationRequest) Validate() error {
	return request.ParseRequest(r)
}

// CreateOrganisationResponse represents the response payload after organization creation.
type CreateOrganisationResponse struct {
	// Unique identifier of the organisation
	ID string `json:"id"`
	// Name of the organisation
	Name string `json:"name"`
	// Timestamp of organisation creation
	CreatedAt time.Time `json:"created_at"`
	// ID of the actor who created the organisation
	CreatedBy uuid.UUID `json:"created_by"`
	// Timestamp of the last update
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// ID of the actor who last updated or created the organisation
	UpdatedBy *uuid.UUID `json:"updated_by,omitempty"`
	// The object version number
	Version int32 `json:"version"`
}
