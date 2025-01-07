package dto

import "encore.app/apicore/common/pipes/request"

// RefreshSessionRequest represents the request payload for session refresh.
type RefreshSessionRequest struct {
	// JWT used to obtain new access tokens
	RefreshToken string `json:"refresh_token" validate:"required" encore:"sensitive"`
}

// Validate implements the validation for RefreshSessionRequest
func (r *RefreshSessionRequest) Validate() error {
	return request.ParseRequest(r)
}

// RefreshSessionResponse represents the response payload after session refresh.
type RefreshSessionResponse struct {
	// A new JWT for accessing protected resources
	AccessToken string `json:"access_token"`
	// A new JWT used to obtain future access tokens
	RefreshToken string `json:"refresh_token"`
}
