package commands

import (
	"github.com/google/uuid"
)

// CreateAccountCommand represents the command to be used to create a new account.
//
// E.g., when a user registers on cursor-demo.
type CreateAccountCommand struct {
	ActorID  *uuid.UUID
	Email    string
	Password string
}

// CreateAccountResult represents the result of a successfully executed CreateAccountCommand
type CreateAccountResult struct {
	AccountID uuid.UUID
}
