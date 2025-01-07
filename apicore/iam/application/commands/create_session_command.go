package commands

import (
	"github.com/google/uuid"
)

// CreateSessionCommand represents the command to create a new session.
//
// E.g., when a user registers or signs-in on cursor-demo.
type CreateSessionCommand struct {
	ActorID  *uuid.UUID
	Email    string
	Password string
}

// CreateSessionResult represents the result of a successful session creation
type CreateSessionResult struct {
	AccessToken  string
	RefreshToken string
}
