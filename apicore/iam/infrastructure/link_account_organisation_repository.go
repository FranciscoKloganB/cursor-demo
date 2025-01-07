package infrastructure

import (
	"context"

	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/iam/application/ports"
	"encore.app/apicore/iam/domain"
	"encore.dev/beta/errs"
	"encore.dev/rlog"
)

// LinkAccountOrganisationRepository implements the ILinkAccountOrganisationRepository interface
type LinkAccountOrganisationRepository struct{}

// NewLinkAccountOrganisationRepository creates a new instance of LinkAccountOrganisationRepository
func NewLinkAccountOrganisationRepository() ports.ILinkAccountOrganisationRepository {
	return &LinkAccountOrganisationRepository{}
}

// Save persists the account-organisation relationship with its role
func (r *LinkAccountOrganisationRepository) Save(ctx context.Context, qrs *db.Queries, accountOrg *domain.AccountOrganisationAggregate) (interface{}, error) {
	params := db.InsertAccountOrganisationRoleParams{
		AccountID:      accountOrg.GetAccountID(),
		OrganisationID: accountOrg.GetOrganisationID(),
		RoleID:         *accountOrg.GetRoleID(),
		CreatedAt:      accountOrg.GetCreatedAt(),
		CreatedBy:      accountOrg.GetCreatedBy(),
		Version:        accountOrg.GetVersion(),
	}

	err := qrs.InsertAccountOrganisationRole(ctx, params)

	if err != nil {
		rlog.Error(
			"Encountered unexpected error while inserting account organisation role link",
			"error", err,
			"accountID", accountOrg.GetAccountID(),
			"organisationID", accountOrg.GetOrganisationID(),
		)

		return nil, errs.WrapCode(err, errs.Internal, "account_organisation_not_linked")
	}

	return nil, nil
}
