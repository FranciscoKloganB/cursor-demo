package commands

import "github.com/google/uuid"

// CreateOrganisationCommand represents the command to create a new organisation.
type CreateOrganisationCommand struct {
	ActorID *uuid.UUID
	Name    string
}
