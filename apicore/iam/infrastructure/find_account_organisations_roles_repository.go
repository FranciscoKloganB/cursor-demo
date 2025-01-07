package infrastructure

import (
	"context"
	"database/sql"

	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/iam/application/ports"
	"encore.app/apicore/iam/domain"
	"encore.app/apicore/iam/domain/entities"
	"encore.app/apicore/iam/domain/valueobjects"
	"encore.dev/beta/errs"
	"encore.dev/rlog"
	"github.com/google/uuid"
)

// FindAccountOrganisationsRolesRepository implements the IFindAccountOrganisationsRolesRepository interface
type FindAccountOrganisationsRolesRepository struct{}

// NewFindAccountOrganisationsRolesRepository creates a new instance of FindAccountOrganisationsRolesRepository
func NewFindAccountOrganisationsRolesRepository() ports.IFindAccountOrganisationsRolesRepository {
	return &FindAccountOrganisationsRolesRepository{}
}

// ByAccountID finds an account and determinates what organisations the account belongs to as well as the role of the account in the each organisation
func (r *FindAccountOrganisationsRolesRepository) ByAccountID(ctx context.Context, qrs *db.Queries, accountID uuid.UUID) (*domain.AccountOrganisationsAggregate, error) {
	rows, err := qrs.FindAccountOrganisationsRoles(ctx, accountID)

	if err != nil {
		if err == sql.ErrNoRows {
			rlog.Info("Account not found", "accountID", accountID)
			return nil, errs.WrapCode(err, errs.NotFound, "account_not_found")
		}

		rlog.Error("Failed to retrieve account organisation roles", "error", err, "accountID", accountID)
		return nil, errs.WrapCode(err, errs.Internal, "account_retrieval_failed")
	}

	if len(rows) == 0 {
		rlog.Info("No roles found for account", "accountID", accountID)
		return nil, errs.WrapCode(err, errs.NotFound, "account_not_found")
	}

	// Create account from first row (all rows will have same account data)
	firstRow := rows[0]
	account, err := entities.NewAccount(
		firstRow.ID,
		firstRow.Email,
		firstRow.PasswordHash,
		firstRow.VerificationStatus,
		firstRow.CreatedAt,
		firstRow.CreatedBy,
		firstRow.DeletedAt,
		firstRow.DeletedBy,
		firstRow.UpdatedAt,
		firstRow.UpdatedBy,
		firstRow.Version,
	)
	if err != nil {
		rlog.Error("Failed to convert persistence model to domain account", "error", err, "accountID", accountID)
		return nil, errs.WrapCode(err, errs.Internal, "account_conversion_failed")
	}

	var roles []valueobjects.AccountRoleVO
	for _, row := range rows {
		if row.OrganisationID != nil && row.RoleSlug != nil {
			role := valueobjects.NewAccountRoleVO(
				*row.OrganisationID,
				*row.RoleSlug,
			)

			roles = append(roles, role)
		}
	}

	return domain.NewAccountOrganisationsAggregate(account, roles), nil
}
