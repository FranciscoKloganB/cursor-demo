package iam

import (
	"encore.app/apicore/common/configs"
	"encore.app/apicore/common/datasource"
	"encore.app/apicore/iam/application"
	"encore.app/apicore/iam/infrastructure"
	"encore.app/apicore/iam/presentation"
	sharedapplication "encore.app/apicore/shared/application"
	"encore.app/utils/modules"
	"encore.dev/rlog"
)

// Options defines key value pairs required to configure IamModule.
type Options struct {
	PrivateKey string
	PublicKey  string
}

// Import defines external modules that IamModule needs to have access to.
type Import struct{}

// Provide defines dependencies that IamModule that are owned or declared within IAM module source code.
type Provide struct {
	Datasource datasource.IDatasource
}

// Inject defines the IAM module configuration
type Inject = modules.ModuleInjector[Import, Provide, Options]

// IIamModule defines the interface for the IAM module.
type IIamModule interface {
	Facade() application.IIamFacade
	Controller() presentation.IIamController
}

// Module implements the IIamModule interface.
type Module struct {
	controller presentation.IIamController
	facade     application.IIamFacade
}

// New initializes and returns a new Module.
func New(i Inject) (*Module, error) {
	rlog.Info("IAM module initializing")

	createAccountRepository := infrastructure.NewCreateAccountRepository()
	createRefreshTokenRepository := infrastructure.NewCreateRefreshTokenRepository()
	findAccountRepository := infrastructure.NewFindAccountRepository()
	findAccountOrganisationsRolesRepository := infrastructure.NewFindAccountOrganisationsRolesRepository()
	findRoleRepository := infrastructure.NewFindRoleRepository()
	linkAccountOrgRepository := infrastructure.NewLinkAccountOrganisationRepository()

	rsaJwtService, err := application.NewRsaJwtService(i.Options.PrivateKey, i.Options.PublicKey)
	if err != nil {
		return nil, err
	}

	// Initialize hashing service with desired options
	hashingService := sharedapplication.NewBcryptHashingService(
		sharedapplication.WithMinLength(configs.PasswordMinLength),
		sharedapplication.WithMinEntropy(70),
		sharedapplication.WithBcryptCost(12),
	)

	facade := application.NewIamFacade(
		i.Provide.Datasource,
		createAccountRepository,
		createRefreshTokenRepository,
		findAccountRepository,
		findAccountOrganisationsRolesRepository,
		findRoleRepository,
		linkAccountOrgRepository,
		hashingService,
		rsaJwtService,
	)

	controller := presentation.NewIamController(facade)

	module := &Module{
		controller: controller,
		facade:     facade,
	}

	rlog.Info("IAM module initialized")

	return module, nil
}

// Facade returns the IAM facade.
func (m *Module) Facade() application.IIamFacade {
	return m.facade
}

// Controller returns the IAM controller.
func (m *Module) Controller() presentation.IIamController {
	return m.controller
}
