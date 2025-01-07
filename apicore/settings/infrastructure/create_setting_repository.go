package infrastructure

import (
	"context"

	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/settings/domain/entities"
	"encore.dev/beta/errs"
	"encore.dev/rlog"
)

// CreateSettingRepository handles the persistence of new setting.
type CreateSettingRepository struct{}

// NewCreateSettingRepository creates a new instance of CreateSettingRepository.
func NewCreateSettingRepository() *CreateSettingRepository {
	return &CreateSettingRepository{}
}

// Save persists a new setting flag to the database.
func (r *CreateSettingRepository) Save(ctx context.Context, qrs *db.Queries, setting entities.Setting) (*entities.Setting, error) {
	params, err := toPersistence(setting)
	if err != nil {
		rlog.Error("Failed to convert setting to persistence model", "error", err)
		return nil, errs.WrapCode(err, errs.Internal, "setting_conversion_failed")
	}

	ent, err := qrs.InsertSetting(ctx, db.InsertSettingParams{
		ID:        params.ID,
		Name:      params.Name,
		Slug:      params.Slug,
		Hint:      params.Hint,
		IsActive:  params.IsActive,
		CreatedAt: params.CreatedAt,
		CreatedBy: params.CreatedBy,
		Version:   params.Version,
	})

	if err != nil {
		rlog.Error("Failed to save setting", "error", err, "settingID", params.ID)
		return nil, errs.WrapCode(err, errs.Internal, "setting_not_saved")
	}

	setting, err = toDomain(ent)
	if err != nil {
		rlog.Error("Failed to convert persistence model to domain setting", "error", err)
		return nil, errs.WrapCode(err, errs.Internal, "setting_conversion_failed")
	}

	return &setting, nil
}
