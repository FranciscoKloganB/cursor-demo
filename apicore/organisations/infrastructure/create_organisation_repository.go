package infrastructure

import (
	"context"

	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/organisations/domain/entities"
	"encore.dev/beta/errs"
	"encore.dev/rlog"
)

// CreateOrganisationRepository handles the persistence of new organisation.
type CreateOrganisationRepository struct{}

// NewCreateOrganisationRepository creates a new instance of CreateOrganisationRepository.
func NewCreateOrganisationRepository() *CreateOrganisationRepository {
	return &CreateOrganisationRepository{}
}

// Save persists a new organisation to the database.
func (r *CreateOrganisationRepository) Save(ctx context.Context, qrs *db.Queries, organisation entities.Organisation) (*entities.Organisation, error) {
	params, err := toPersistence(organisation)
	if err != nil {
		rlog.Error("Failed to convert organisation to persistence model", "error", err)
		return nil, errs.WrapCode(err, errs.Internal, "organisation_conversion_failed")
	}

	ent, err := qrs.InsertOrganisation(ctx, db.InsertOrganisationParams{
		ID:        params.ID,
		Name:      params.Name,
		CreatedAt: params.CreatedAt,
		CreatedBy: params.CreatedBy,
		Version:   params.Version,
	})

	if err != nil {
		rlog.Error("Failed to save organisation", "error", err, "organisationID", params.ID)
		return nil, errs.WrapCode(err, errs.Internal, "organisation_not_saved")
	}

	organisation, err = toDomain(ent)
	if err != nil {
		rlog.Error("Failed to convert persistence model to domain organisation", "error", err)
		return nil, errs.WrapCode(err, errs.Internal, "organisation_conversion_failed")
	}

	return &organisation, nil
}
