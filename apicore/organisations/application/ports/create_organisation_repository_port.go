package ports

import (
	"context"

	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/organisations/domain/entities"
)

// ICreateOrganisationRepository defines the interface for the organisation repository.
type ICreateOrganisationRepository interface {
	Save(ctx context.Context, qrs *db.Queries, organisation entities.Organisation) (*entities.Organisation, error)
}
