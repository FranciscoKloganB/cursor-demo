package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Account represents an account in the domain layer.
type Account struct {
	ID                 uuid.UUID
	Email              string
	Password           *string
	VerificationStatus string
	CreatedAt          time.Time
	CreatedBy          uuid.UUID
	DeletedAt          *time.Time
	DeletedBy          *uuid.UUID
	UpdatedAt          *time.Time
	UpdatedBy          *uuid.UUID
	Version            int32
}

// NewAccount creates a new account instance.
func NewAccount(
	id uuid.UUID,
	email string,
	password string,
	verificationStatus string,
	createdAt time.Time,
	createdBy uuid.UUID,
	deletedAt *time.Time,
	deletedBy *uuid.UUID,
	updatedAt *time.Time,
	updatedBy *uuid.UUID,
	version int32,
) (Account, error) {
	if email == "" {
		return Account{}, errors.New("email cannot be empty")
	}
	if password == "" {
		return Account{}, errors.New("password hash cannot be empty")
	}
	if verificationStatus == "" {
		return Account{}, errors.New("verification status cannot be empty")
	}

	return Account{
		ID:                 id,
		Email:              email,
		Password:           &password,
		VerificationStatus: verificationStatus,
		CreatedAt:          createdAt,
		CreatedBy:          createdBy,
		DeletedAt:          deletedAt,
		DeletedBy:          deletedBy,
		UpdatedAt:          updatedAt,
		UpdatedBy:          updatedBy,
		Version:            version,
	}, nil
}
