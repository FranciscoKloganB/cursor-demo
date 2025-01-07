package settings

import (
	"encore.app/apicore/common/datasource"
	"encore.app/apicore/settings/application"
	"encore.app/apicore/settings/infrastructure"
	"encore.app/apicore/settings/presentation"
	"encore.app/utils/modules"
	"encore.dev/rlog"
)

// Options defines key value pairs required to configure SettingsModule.
type Options struct{}

// Import defines external modules that SettingsModule needs to have access to.
type Import struct{}

// Provide defines dependencies that SettingsModule that are owned or declared within Settings module source code.
type Provide struct {
	Datasource datasource.IDatasource
}

// Inject defines the Settings module configuration
type Inject = modules.ModuleInjector[Import, Provide, Options]

// ISettingsModule defines the interface for the settings module.
type ISettingsModule interface {
	Facade() application.ISettingsFacade
	Controller() presentation.ISettingsController
}

// Module implements the ISettingsModule interface.
type Module struct {
	controller presentation.ISettingsController
	facade     application.ISettingsFacade
}

// New initializes and returns a new Module.
func New(i Inject) (*Module, error) {
	rlog.Info("Settings module initializing")

	createSettingRepository := infrastructure.NewCreateSettingRepository()
	findSettingRepository := infrastructure.NewFindSettingRepository()

	facade := application.NewSettingsFacade(
		i.Provide.Datasource,
		createSettingRepository,
		findSettingRepository,
	)
	controller := presentation.NewSettingsController(facade)

	module := &Module{
		controller: controller,
		facade:     facade,
	}

	rlog.Info("Settings module initialized")

	return module, nil
}

// Facade returns the settings facade.
func (m *Module) Facade() application.ISettingsFacade {
	return m.facade
}

// Controller returns the settings controller.
func (m *Module) Controller() presentation.ISettingsController {
	return m.controller
}
