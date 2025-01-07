package presentation

import (
	"context"

	"encore.app/apicore/organisations/application"
	"encore.app/apicore/organisations/application/commands"
	"encore.app/apicore/organisations/presentation/dto"
	"encore.app/utils"
	"encore.dev/rlog"
)

// IOrganisationsController defines the interface for the organisations controller.
type IOrganisationsController interface {
	CreateOrganisation(ctx context.Context, params *dto.CreateOrganisationRequest) (*dto.CreateOrganisationResponse, error)
}

// OrganisationsController handles incomming HTTP requests and forwards them to the application.
type OrganisationsController struct {
	organisationsFacade application.IOrganisationsFacade
}

// NewOrganisationsController creates a new instance of OrganisationsController.
func NewOrganisationsController(organisationsFacade application.IOrganisationsFacade) *OrganisationsController {
	return &OrganisationsController{
		organisationsFacade: organisationsFacade,
	}
}

// CreateOrganisation handles HTTP requests to create an Organisation.
func (c *OrganisationsController) CreateOrganisation(ctx context.Context, params *dto.CreateOrganisationRequest) (*dto.CreateOrganisationResponse, error) {
	cmd := commands.CreateOrganisationCommand{
		ActorID: utils.GetCtxActor(ctx),
		Name:    params.Name,
	}

	organisation, err := c.organisationsFacade.CreateOrganisation(ctx, cmd)

	if err != nil {
		rlog.Error("Failed to create organisation", "error", err, "organisationName", params.Name)
		return nil, err
	}

	response := &dto.CreateOrganisationResponse{
		ID:        organisation.ID.String(),
		Name:      organisation.Name,
		CreatedAt: organisation.CreatedAt,
		CreatedBy: organisation.CreatedBy,
		UpdatedAt: organisation.UpdatedAt,
		UpdatedBy: organisation.UpdatedBy,
		Version:   organisation.Version,
	}

	return response, nil
}
