package organisations

import (
	"encore.app/apicore/common/datasource"
	"encore.app/apicore/organisations/application"
	"encore.app/apicore/organisations/infrastructure"
	"encore.app/apicore/organisations/presentation"
	"encore.app/utils/modules"
	"encore.dev/rlog"
)

// Options defines key value pairs required to configure OrganisationsModule.
type Options struct{}

// Import defines external modules that OrganisationsModule needs to have access to.
type Import struct{}

// Provide defines dependencies that OrganisationsModule that are owned or declared within Organisations module source code.
type Provide struct {
	Datasource datasource.IDatasource
}

// Inject defines the Organisations module configuration
type Inject = modules.ModuleInjector[Import, Provide, Options]

// IOrganisationsModule defines the interface for the organisations module.
type IOrganisationsModule interface {
	Facade() application.IOrganisationsFacade
	Controller() presentation.IOrganisationsController
}

// Module implements the IOrganisationsModule interface.
type Module struct {
	controller presentation.IOrganisationsController
	facade     application.IOrganisationsFacade
}

// New initializes and returns a new Module.
func New(i Inject) (*Module, error) {
	rlog.Info("Organisations module initializing")

	createOrganisationRepository := infrastructure.NewCreateOrganisationRepository()
	facade := application.NewOrganisationsFacade(i.Provide.Datasource, createOrganisationRepository)
	controller := presentation.NewOrganisationsController(facade)

	module := &Module{
		controller: controller,
		facade:     facade,
	}

	rlog.Info("Organisations module initialized")

	return module, nil
}

// Facade returns the organisations facade.
func (m *Module) Facade() application.IOrganisationsFacade {
	return m.facade
}

// Controller returns the organisations controller.
func (m *Module) Controller() presentation.IOrganisationsController {
	return m.controller
}
