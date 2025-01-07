package application

import (
	"context"

	"encore.app/apicore/common/datasource"
	"encore.app/apicore/settings/application/commands"
	"encore.app/apicore/settings/application/ports"
	"encore.app/apicore/settings/domain/entities"
	"encore.dev/rlog"
)

// ISettingsFacade defines the interface for the settings facade.
type ISettingsFacade interface {
	CreateSetting(ctx context.Context, cmd commands.CreateSettingCommand) (entities.Setting, error)
	GetSettingByID(ctx context.Context, id string) (entities.Setting, error)
}

// SettingsFacade provides methods to manage settings.
type SettingsFacade struct {
	datasource              datasource.IDatasource
	createSettingRepository ports.ICreateSettingRepository
	findSettingRepository   ports.IFindSettingRepository
}

// NewSettingsFacade creates a new instance of SettingsFacade.
func NewSettingsFacade(
	datasrc datasource.IDatasource,
	createSettingRepository ports.ICreateSettingRepository,
	findSettingRepository ports.IFindSettingRepository,
) *SettingsFacade {
	return &SettingsFacade{
		datasource:              datasrc,
		createSettingRepository: createSettingRepository,
		findSettingRepository:   findSettingRepository,
	}
}

// CreateSetting handles the creation of a new setting.
func (f *SettingsFacade) CreateSetting(ctx context.Context, cmd commands.CreateSettingCommand) (entities.Setting, error) {
	rlog.Info("Creating new setting", "name", cmd.Name)

	setting, err := createSettingFactory(cmd)
	if err != nil {
		rlog.Error("Failed to create setting from command", "error", err, "name", cmd.Name)
		return setting, err
	}

	rlog.Debug("Setting created from factory", "setting_id", setting.ID)

	createdSetting, err := f.createSettingRepository.Save(ctx, f.datasource.Queries(), setting)
	if err != nil {
		rlog.Error("Failed to persist setting", "error", err, "setting_id", setting.ID)
		return entities.Setting{}, err
	}

	rlog.Info("Setting created successfully",
		"setting_id", createdSetting.ID,
		"name", createdSetting.Name,
		"slug", createdSetting.Slug,
	)

	return *createdSetting, nil
}

// GetSettingByID retrieves a setting by its ID
func (f *SettingsFacade) GetSettingByID(ctx context.Context, id string) (entities.Setting, error) {
	rlog.Info("Retrieving setting", "id", id)

	setting, err := f.findSettingRepository.ByID(ctx, f.datasource.Queries(), id)
	if err != nil {
		rlog.Error("Failed to retrieve setting", "error", err, "id", id)
		return entities.Setting{}, err
	}

	return *setting, nil
}
