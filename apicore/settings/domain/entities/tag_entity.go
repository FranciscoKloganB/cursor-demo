package entities

import (
	"time"

	"github.com/google/uuid"
)

// Tag represents a tag associated with a setting flag.
type Tag struct {
	ID        string
	Name      string
	Color     string
	CreatedAt time.Time
	CreatedBy uuid.UUID
	DeletedAt *time.Time
	DeletedBy *uuid.UUID
	UpdatedAt *time.Time
	UpdatedBy *uuid.UUID
	Version   int32
}
