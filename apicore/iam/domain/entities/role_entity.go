package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Role represents a role in the domain layer.
type Role struct {
	ID        *int32
	Slug      string
	CreatedAt time.Time
	CreatedBy uuid.UUID
	DeletedAt *time.Time
	DeletedBy *uuid.UUID
	UpdatedAt *time.Time
	UpdatedBy *uuid.UUID
	Version   int32
}

// NewRole creates a new role instance.
func NewRole(
	id *int32,
	slug string,
	createdAt time.Time,
	createdBy uuid.UUID,
	deletedAt *time.Time,
	deletedBy *uuid.UUID,
	updatedAt *time.Time,
	updatedBy *uuid.UUID,
	version int32,
) (Role, error) {
	if slug == "" {
		return Role{}, errors.New("role slug cannot be empty")
	}

	return Role{
		ID:        id,
		Slug:      slug,
		CreatedAt: createdAt,
		CreatedBy: createdBy,
		DeletedAt: deletedAt,
		DeletedBy: deletedBy,
		UpdatedAt: updatedAt,
		UpdatedBy: updatedBy,
		Version:   version,
	}, nil
}
