package commands

import (
	"time"

	"github.com/google/uuid"
)

// CreateRefreshTokenCommand represents the command to create a new refresh token.
//
// E.g., when a user session is created or refreshed.
type CreateRefreshTokenCommand struct {
	AccountID  uuid.UUID
	TimeToLive time.Duration
	TokenValue string
}
