package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Organisation represents an organisation in the domain layer.
type Organisation struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	CreatedBy uuid.UUID
	DeletedAt *time.Time
	DeletedBy *uuid.UUID
	UpdatedAt *time.Time
	UpdatedBy *uuid.UUID
	Version   int32
}

// NewOrganisation creates a new organisation instance.
func NewOrganisation(id uuid.UUID, name string, createdAt time.Time, createdBy uuid.UUID, deletedAt *time.Time, deletedBy *uuid.UUID, updatedAt *time.Time, updatedBy *uuid.UUID, version int32) (Organisation, error) {
	if name == "" {
		return Organisation{}, errors.New("organisation name cannot be empty")
	}

	domainEntity := Organisation{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
		CreatedBy: createdBy,
		DeletedAt: deletedAt,
		DeletedBy: deletedBy,
		UpdatedAt: updatedAt,
		UpdatedBy: updatedBy,
		Version:   version,
	}

	return domainEntity, nil
}
