// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID                 uuid.UUID  `db:"id" json:"id"`
	Email              string     `db:"email" json:"email"`
	PasswordHash       string     `db:"password_hash" json:"password_hash"`
	VerificationStatus string     `db:"verification_status" json:"verification_status"`
	CreatedAt          time.Time  `db:"created_at" json:"created_at"`
	CreatedBy          uuid.UUID  `db:"created_by" json:"created_by"`
	DeletedAt          *time.Time `db:"deleted_at" json:"deleted_at"`
	DeletedBy          *uuid.UUID `db:"deleted_by" json:"deleted_by"`
	UpdatedAt          *time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy          *uuid.UUID `db:"updated_by" json:"updated_by"`
	Version            int32      `db:"version" json:"version"`
}

type AccountOrganisationRole struct {
	AccountID      uuid.UUID  `db:"account_id" json:"account_id"`
	OrganisationID uuid.UUID  `db:"organisation_id" json:"organisation_id"`
	RoleID         int32      `db:"role_id" json:"role_id"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	CreatedBy      uuid.UUID  `db:"created_by" json:"created_by"`
	DeletedAt      *time.Time `db:"deleted_at" json:"deleted_at"`
	DeletedBy      *uuid.UUID `db:"deleted_by" json:"deleted_by"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy      *uuid.UUID `db:"updated_by" json:"updated_by"`
	Version        int32      `db:"version" json:"version"`
}

type AccountOrganisationScope struct {
	AccountID      uuid.UUID  `db:"account_id" json:"account_id"`
	OrganisationID uuid.UUID  `db:"organisation_id" json:"organisation_id"`
	ScopeID        int32      `db:"scope_id" json:"scope_id"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	CreatedBy      uuid.UUID  `db:"created_by" json:"created_by"`
	DeletedAt      *time.Time `db:"deleted_at" json:"deleted_at"`
	DeletedBy      *uuid.UUID `db:"deleted_by" json:"deleted_by"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy      *uuid.UUID `db:"updated_by" json:"updated_by"`
	Version        int32      `db:"version" json:"version"`
}

type AccountsOrganisation struct {
	AccountID      uuid.UUID  `db:"account_id" json:"account_id"`
	OrganisationID uuid.UUID  `db:"organisation_id" json:"organisation_id"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	CreatedBy      uuid.UUID  `db:"created_by" json:"created_by"`
	DeletedAt      *time.Time `db:"deleted_at" json:"deleted_at"`
	DeletedBy      *uuid.UUID `db:"deleted_by" json:"deleted_by"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy      *uuid.UUID `db:"updated_by" json:"updated_by"`
	Version        int32      `db:"version" json:"version"`
}

type OneTimePassword struct {
	ID        int32      `db:"id" json:"id"`
	AccountID *uuid.UUID `db:"account_id" json:"account_id"`
	ExpiresAt time.Time  `db:"expires_at" json:"expires_at"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	CreatedBy uuid.UUID  `db:"created_by" json:"created_by"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
	DeletedBy *uuid.UUID `db:"deleted_by" json:"deleted_by"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy *uuid.UUID `db:"updated_by" json:"updated_by"`
	Version   int32      `db:"version" json:"version"`
}

type Organisation struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	CreatedBy uuid.UUID  `db:"created_by" json:"created_by"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
	DeletedBy *uuid.UUID `db:"deleted_by" json:"deleted_by"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy *uuid.UUID `db:"updated_by" json:"updated_by"`
	Version   int32      `db:"version" json:"version"`
}

type RefreshToken struct {
	ID         uuid.UUID  `db:"id" json:"id"`
	AccountID  *uuid.UUID `db:"account_id" json:"account_id"`
	IsRevoked  bool       `db:"is_revoked" json:"is_revoked"`
	TokenValue string     `db:"token_value" json:"token_value"`
	ExpiresAt  time.Time  `db:"expires_at" json:"expires_at"`
	LastUsedAt *time.Time `db:"last_used_at" json:"last_used_at"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	CreatedBy  uuid.UUID  `db:"created_by" json:"created_by"`
	DeletedAt  *time.Time `db:"deleted_at" json:"deleted_at"`
	DeletedBy  *uuid.UUID `db:"deleted_by" json:"deleted_by"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy  *uuid.UUID `db:"updated_by" json:"updated_by"`
	Version    int32      `db:"version" json:"version"`
}

type Role struct {
	ID        int32      `db:"id" json:"id"`
	Slug      string     `db:"slug" json:"slug"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	CreatedBy uuid.UUID  `db:"created_by" json:"created_by"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
	DeletedBy *uuid.UUID `db:"deleted_by" json:"deleted_by"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy *uuid.UUID `db:"updated_by" json:"updated_by"`
	Version   int32      `db:"version" json:"version"`
}

type Scope struct {
	ID        int32      `db:"id" json:"id"`
	Slug      string     `db:"slug" json:"slug"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	CreatedBy uuid.UUID  `db:"created_by" json:"created_by"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
	DeletedBy *uuid.UUID `db:"deleted_by" json:"deleted_by"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy *uuid.UUID `db:"updated_by" json:"updated_by"`
	Version   int32      `db:"version" json:"version"`
}

type Setting struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	Slug      string     `db:"slug" json:"slug"`
	Hint      string     `db:"hint" json:"hint"`
	IsActive  bool       `db:"is_active" json:"is_active"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	CreatedBy uuid.UUID  `db:"created_by" json:"created_by"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
	DeletedBy *uuid.UUID `db:"deleted_by" json:"deleted_by"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy *uuid.UUID `db:"updated_by" json:"updated_by"`
	Version   int32      `db:"version" json:"version"`
}
