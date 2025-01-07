package commands

import "github.com/google/uuid"

// CreateSettingCommand represents the command to create a new setting.
type CreateSettingCommand struct {
	ActorID  *uuid.UUID
	Name     string
	Slug     string
	Hint     string
	IsActive bool
}
