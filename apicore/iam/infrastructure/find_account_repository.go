package infrastructure

import (
	"context"
	"database/sql"
	"errors"

	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/iam/application/ports"
	"encore.app/apicore/iam/domain/entities"
	"encore.dev/beta/errs"
	"encore.dev/rlog"
)

// FindAccountRepository implements the IFindAccountRepository interface
type FindAccountRepository struct{}

// NewFindAccountRepository creates a new instance of FindAccountRepository
func NewFindAccountRepository() ports.IFindAccountRepository {
	return &FindAccountRepository{}
}

// ByEmail finds an account by its email
func (r *FindAccountRepository) ByEmail(ctx context.Context, qrs *db.Queries, email string) (*entities.Account, error) {
	accountEntity, err := qrs.FindAccountByEmail(ctx, email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			rlog.Info("Account not found", "email", email)
			return nil, errs.WrapCode(err, errs.NotFound, "account_not_found")
		}

		rlog.Error(
			"Encountered unexpected error while querying accounts table by email",
			"error", err,
			"email", email,
		)

		return nil, errs.WrapCode(err, errs.Internal, "account_not_retrieved")
	}

	account, err := toDomain(accountEntity)
	if err != nil {
		rlog.Error("Failed to map account entity to domain", "error", err)
		return nil, errs.WrapCode(err, errs.Internal, "account_mapping_failed")
	}

	return &account, nil
}
