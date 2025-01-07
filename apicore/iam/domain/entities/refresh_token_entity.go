package entities

import (
	"time"

	"encore.dev/beta/errs"
	"github.com/google/uuid"
)

// RefreshToken represents a refresh token in the domain layer.
type RefreshToken struct {
	ID         uuid.UUID
	AccountID  uuid.UUID
	IsRevoked  bool
	Value      string
	ExpiresAt  time.Time
	LastUsedAt *time.Time
	CreatedAt  time.Time
	CreatedBy  uuid.UUID
	DeletedAt  *time.Time
	DeletedBy  *uuid.UUID
	UpdatedAt  *time.Time
	UpdatedBy  *uuid.UUID
	Version    int32
}

// NewRefreshToken creates a new refresh token instance.
func NewRefreshToken(
	id uuid.UUID,
	accountID uuid.UUID,
	tokenValue string,
	expiresAt time.Time,
	createdAt time.Time,
	createdBy uuid.UUID,
	deletedAt *time.Time,
	deletedBy *uuid.UUID,
	updatedAt *time.Time,
	updatedBy *uuid.UUID,
	version int32,
) (RefreshToken, error) {
	if id == uuid.Nil {
		return RefreshToken{}, &errs.Error{
			Code:    errs.Internal,
			Message: "refresh_token_entity_invalid",
			Meta: errs.Metadata{
				"Issue": "id cannot be nil",
			},
		}
	}
	if accountID == uuid.Nil {
		return RefreshToken{}, &errs.Error{
			Code:    errs.Internal,
			Message: "refresh_token_entity_invalid",
			Meta: errs.Metadata{
				"Issue": "accountID cannot be nil",
			},
		}
	}
	if tokenValue == "" {
		return RefreshToken{}, &errs.Error{
			Code:    errs.Internal,
			Message: "refresh_token_entity_invalid",
			Meta: errs.Metadata{
				"Issue": "token can not be empty",
			},
		}
	}

	if expiresAt.IsZero() {
		return RefreshToken{}, &errs.Error{
			Code:    errs.Internal,
			Message: "refresh_token_entity_invalid",
			Meta: errs.Metadata{
				"Issue": "expiresAt can not be empty",
			},
		}
	}

	return RefreshToken{
		ID:         id,
		AccountID:  accountID,
		IsRevoked:  false,
		Value:      tokenValue,
		ExpiresAt:  expiresAt,
		LastUsedAt: nil,
		CreatedAt:  createdAt,
		CreatedBy:  createdBy,
		DeletedAt:  deletedAt,
		DeletedBy:  deletedBy,
		UpdatedAt:  updatedAt,
		UpdatedBy:  updatedBy,
		Version:    version,
	}, nil
}

// SetTimeToLive updates the expiration time of the refresh token
// by adding the specified time-to-live duration to the current time.
func (t *RefreshToken) SetTimeToLive(ttl time.Duration) {
	t.ExpiresAt = time.Now().Add(ttl).UTC()
}
