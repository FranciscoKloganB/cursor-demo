package ports

import (
	"context"

	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/iam/domain"
	"github.com/google/uuid"
)

// IFindAccountOrganisationsRolesRepository defines methods for finding an account,
// the organisations the account belongs as well as the roles of the account in
// that organisation.
type IFindAccountOrganisationsRolesRepository interface {
	// ByAccountID finds an account and related organisation data by the account id
	ByAccountID(ctx context.Context, qrs *db.Queries, accountID uuid.UUID) (*domain.AccountOrganisationsAggregate, error)
}
