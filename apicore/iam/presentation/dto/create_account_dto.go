package dto

import (
	"encore.app/apicore/common/pipes/request"
)

// CreateAccountRequest represents the request payload for account creation.
type CreateAccountRequest struct {
	// Email address that will be used to create new sessions for the account
	Email string `json:"email" mod:"trim" validate:"required,email" encore:"sensitive"`
	// Password for the account, must be at least 16 characters
	Password string `json:"password" validate:"required,min=12" encore:"sensitive"`
}

// Validate implements the validation for CreateAccountRequest
func (r *CreateAccountRequest) Validate() error {
	return request.ParseRequest(r)
}

// CreateAccountResponse represents the response payload after account creation.
type CreateAccountResponse struct {
	// Unique identifier of the created account
	AccountID string `json:"account_id"`
}
