package infrastructure

import (
	"context"

	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/iam/domain/entities"
	"encore.dev/beta/errs"
	"encore.dev/rlog"
)

// CreateAccountRepository handles the persistence of new accounts.
type CreateAccountRepository struct{}

// NewCreateAccountRepository creates a new instance of CreateAccountRepository.
func NewCreateAccountRepository() *CreateAccountRepository {
	return &CreateAccountRepository{}
}

// Save persists a new account to the database.
func (r *CreateAccountRepository) Save(ctx context.Context, qrs *db.Queries, account entities.Account) (*entities.Account, error) {
	params, err := toPersistence(account)

	if err != nil {
		rlog.Error("Failed to convert account to persistence model", "error", err)
		return nil, errs.WrapCode(err, errs.Internal, "account_conversion_failed")
	}

	ent, err := qrs.InsertAccount(ctx, db.InsertAccountParams{
		ID:                 params.ID,
		Email:              params.Email,
		PasswordHash:       params.PasswordHash,
		VerificationStatus: params.VerificationStatus,
		CreatedAt:          params.CreatedAt,
		CreatedBy:          params.CreatedBy,
		Version:            params.Version,
	})

	if err != nil {
		rlog.Error("Failed to save account", "error", err, "accountID", params.ID)
		return nil, errs.WrapCode(err, errs.Internal, "account_not_saved")
	}

	account, err = toDomain(ent)
	if err != nil {
		rlog.Error("Failed to convert persistence model to domain account", "error", err)
		return nil, errs.WrapCode(err, errs.Internal, "account_conversion_failed")
	}

	return &account, nil
}
