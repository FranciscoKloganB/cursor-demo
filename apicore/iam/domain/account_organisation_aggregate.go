package domain

import (
	"time"

	"encore.app/apicore/iam/domain/entities"
	"github.com/google/uuid"
)

// AccountOrganisationAggregate represents the relationship between an account and its role in an organisation
type AccountOrganisationAggregate struct {
	account      entities.Account
	role         entities.Role
	organisation uuid.UUID

	// Audit fields for the relationship itself
	createdAt time.Time
	createdBy uuid.UUID
	deletedAt *time.Time
	deletedBy *uuid.UUID
	updatedAt *time.Time
	updatedBy *uuid.UUID
	version   int32
}

// NewAccountOrganisationAggregate creates a new account-organisation relationship
func NewAccountOrganisationAggregate(
	account entities.Account,
	role entities.Role,
	organisationID uuid.UUID,
	createdAt time.Time,
	createdBy uuid.UUID,
	deletedAt *time.Time,
	deletedBy *uuid.UUID,
	updatedAt *time.Time,
	updatedBy *uuid.UUID,
	version int32,
) *AccountOrganisationAggregate {
	return &AccountOrganisationAggregate{
		account:      account,
		role:         role,
		organisation: organisationID,
		createdAt:    createdAt,
		createdBy:    createdBy,
		deletedAt:    deletedAt,
		deletedBy:    deletedBy,
		updatedAt:    updatedAt,
		updatedBy:    updatedBy,
		version:      version,
	}
}

// GetAccountID returns the account ID
func (a *AccountOrganisationAggregate) GetAccountID() uuid.UUID {
	return a.account.ID
}

// GetRoleID returns the role ID
func (a *AccountOrganisationAggregate) GetRoleID() *int32 {
	return a.role.ID
}

// GetOrganisationID returns the organisation ID
func (a *AccountOrganisationAggregate) GetOrganisationID() uuid.UUID {
	return a.organisation
}

// GetCreatedAt returns the creation timestamp
func (a *AccountOrganisationAggregate) GetCreatedAt() time.Time {
	return a.createdAt
}

// GetCreatedBy returns the ID of the creator
func (a *AccountOrganisationAggregate) GetCreatedBy() uuid.UUID {
	return a.createdBy
}

// GetDeletedAt returns the deletion timestamp if any
func (a *AccountOrganisationAggregate) GetDeletedAt() *time.Time {
	return a.deletedAt
}

// GetDeletedBy returns the ID of the user who deleted if any
func (a *AccountOrganisationAggregate) GetDeletedBy() *uuid.UUID {
	return a.deletedBy
}

// GetUpdatedAt returns the last update timestamp if any
func (a *AccountOrganisationAggregate) GetUpdatedAt() *time.Time {
	return a.updatedAt
}

// GetUpdatedBy returns the ID of the last user who updated if any
func (a *AccountOrganisationAggregate) GetUpdatedBy() *uuid.UUID {
	return a.updatedBy
}

// GetVersion returns the version number
func (a *AccountOrganisationAggregate) GetVersion() int32 {
	return a.version
}
