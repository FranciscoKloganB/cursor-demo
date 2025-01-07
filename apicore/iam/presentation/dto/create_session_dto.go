package dto

import "encore.app/apicore/common/pipes/request"

// CreateSessionRequest represents the request payload for session creation.
type CreateSessionRequest struct {
	// Email address associated with the account
	Email string `json:"email" mod:"trim" validate:"required,email" encore:"sensitive"`
	// Account password
	Password string `json:"password" validate:"required" encore:"sensitive"`
}

// Validate implements the validation for CreateSessionRequest
func (r *CreateSessionRequest) Validate() error {
	return request.ParseRequest(r)
}

// CreateSessionResponse represents the response payload after session creation.
type CreateSessionResponse struct {
	// JWT for accessing protected resources
	AccessToken string `json:"access_token"`
	// JWT used to obtain new access tokens
	RefreshToken string `json:"refresh_token"`
}
