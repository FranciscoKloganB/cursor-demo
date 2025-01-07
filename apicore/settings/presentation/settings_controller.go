package presentation

import (
	"context"
	"fmt"

	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
	"encore.dev/rlog"

	"encore.app/apicore/common/pipes/request"
	settingsapplication "encore.app/apicore/settings/application"
	"encore.app/apicore/settings/application/commands"
	"encore.app/apicore/settings/presentation/dto"
	"github.com/google/uuid"
)

// ISettingsController defines the interface for the settings controller.
type ISettingsController interface {
	CreateSetting(ctx context.Context, params *dto.CreateSettingRequest) (*dto.CreateSettingResponse, error)
	GetSettingByID(ctx context.Context, id string) (*dto.GetSettingResponse, error)
}

// SettingsController handles incoming HTTP requests and forwards them to the application.
type SettingsController struct {
	settingsFacade settingsapplication.ISettingsFacade
}

// NewSettingsController creates a new instance of SettingsController.
func NewSettingsController(settingsFacade settingsapplication.ISettingsFacade) *SettingsController {
	return &SettingsController{
		settingsFacade: settingsFacade,
	}
}

// Define a type that implements errs.ErrDetails
type CustomErrDetails struct {
	segment string
}

// Implement the necessary methods for CustomErrDetails to satisfy errs.ErrDetails
func (d CustomErrDetails) ErrDetails() string {
	return d.segment
}

// CreateSetting handles HTTP requests to create a Setting.
func (c *SettingsController) CreateSetting(ctx context.Context, params *dto.CreateSettingRequest) (*dto.CreateSettingResponse, error) {
	// Get the authenticated user ID using Encore's auth package
	userID, ok := auth.UserID()
	if !ok {
		rlog.Error("No authenticated user found")
		return nil, fmt.Errorf("no authenticated user found")
	}

	// Parse the userID to a UUID
	actorID, err := uuid.Parse(string(userID))
	if err != nil {
		rlog.Error("Invalid user ID format", "error", err)
		return nil, fmt.Errorf("invalid user ID format")
	}

	cmd := commands.CreateSettingCommand{
		ActorID:  &actorID, // Use the parsed UUID
		Name:     params.Name,
		Slug:     params.Slug,
		Hint:     params.Hint,
		IsActive: params.IsEnabled,
	}

	setting, err := c.settingsFacade.CreateSetting(ctx, cmd)

	if err != nil {
		rlog.Error("Failed to create setting", "error", err, "settingName", params.Name)
		return nil, err
	}

	response := &dto.CreateSettingResponse{
		ID:        setting.ID.String(),
		Name:      setting.Name,
		Slug:      setting.Slug,
		Hint:      setting.Hint,
		IsEnabled: setting.IsActive,
		CreatedAt: setting.CreatedAt,
		CreatedBy: setting.CreatedBy,
		UpdatedAt: setting.UpdatedAt,
		UpdatedBy: setting.UpdatedBy,
		Version:   setting.Version,
	}

	return response, nil
}

// GetSettingByID handles HTTP requests to retrieve a Setting by ID.
func (c *SettingsController) GetSettingByID(ctx context.Context, id string) (*dto.GetSettingResponse, error) {
	// Validate that the input id is a valid UUID
	settingUid, err := uuid.Parse(id)

	if err != nil || settingUid == uuid.Nil {
		rlog.Error("Invalid ID format", "error", err, "id", id)

		errDetails := request.ValidationErrors{
			{Field: "id", Error: "invalid UUID format"},
		}

		return nil, &errs.Error{
			Code:    errs.InvalidArgument,
			Message: "validation_failed",
			Details: errDetails,
		}
	}

	setting, err := c.settingsFacade.GetSettingByID(ctx, id)
	if err != nil {
		rlog.Error("Failed to retrieve setting", "error", err, "id", id)
		return nil, err
	}

	response := &dto.GetSettingResponse{
		ID:        setting.ID.String(),
		Name:      setting.Name,
		Slug:      setting.Slug,
		Hint:      setting.Hint,
		IsEnabled: setting.IsActive,
		CreatedAt: setting.CreatedAt,
		CreatedBy: setting.CreatedBy,
		UpdatedAt: setting.UpdatedAt,
		UpdatedBy: setting.UpdatedBy,
		Version:   setting.Version,
	}

	return response, nil
}
