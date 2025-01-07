package ports

import (
	"context"

	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/iam/domain"
)

// ILinkAccountOrganisationRepository defines the interface for linking accounts with organisations.
type ILinkAccountOrganisationRepository interface {
	Save(ctx context.Context, qrs *db.Queries, accountOrg *domain.AccountOrganisationAggregate) (interface{}, error)
}
