package application

import (
	"context"

	"encore.app/apicore/common/datasource"
	"encore.app/apicore/organisations/application/commands"
	"encore.app/apicore/organisations/application/ports"
	"encore.app/apicore/organisations/domain/entities"
	"encore.dev/rlog"
)

// IOrganisationsFacade defines the interface for the organisations facade.
type IOrganisationsFacade interface {
	CreateOrganisation(ctx context.Context, cmd commands.CreateOrganisationCommand) (entities.Organisation, error)
}

// OrganisationsFacade provides methods to manage organisations.
type OrganisationsFacade struct {
	datasource                   datasource.IDatasource
	createOrganisationRepository ports.ICreateOrganisationRepository
}

// NewOrganisationsFacade creates a new instance of OrganisationsFacade.
func NewOrganisationsFacade(
	datasrc datasource.IDatasource,
	createOrganisationRepository ports.ICreateOrganisationRepository,
) *OrganisationsFacade {
	return &OrganisationsFacade{
		datasource:                   datasrc,
		createOrganisationRepository: createOrganisationRepository,
	}
}

// CreateOrganisation handles the creation of a new organisation.
func (f *OrganisationsFacade) CreateOrganisation(ctx context.Context, cmd commands.CreateOrganisationCommand) (entities.Organisation, error) {
	rlog.Info("Creating new organisation", "name", cmd.Name)

	organisation, err := createOrganisationFactory(cmd)
	if err != nil {
		rlog.Error("Failed to create organisation from command", "error", err, "name", cmd.Name)
		return organisation, err
	}

	rlog.Debug("Organisation created from factory", "organisation_id", organisation.ID)

	createdOrganisation, err := f.createOrganisationRepository.Save(ctx, f.datasource.Queries(), organisation)
	if err != nil {
		rlog.Error("Failed to persist organisation", "error", err, "organisation_id", organisation.ID)
		return entities.Organisation{}, err
	}

	rlog.Info("Organisation created successfully",
		"organisation_id", createdOrganisation.ID,
		"name", createdOrganisation.Name,
	)

	return *createdOrganisation, nil
}
