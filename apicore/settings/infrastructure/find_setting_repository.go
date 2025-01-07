package infrastructure

import (
	"context"
	"database/sql"
	"errors"

	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/settings/application/ports"
	"encore.app/apicore/settings/domain/entities"
	"encore.dev/beta/errs"
	"encore.dev/rlog"
	"github.com/google/uuid"
)

// FindSettingRepository implements the IFindSettingRepository interface
type FindSettingRepository struct{}

// NewFindSettingRepository creates a new instance of FindSettingRepository
func NewFindSettingRepository() ports.IFindSettingRepository {
	return &FindSettingRepository{}
}

// ByID finds a setting by its ID
func (r *FindSettingRepository) ByID(ctx context.Context, qrs *db.Queries, id string) (*entities.Setting, error) {
	settingID, err := uuid.Parse(id)
	if err != nil {
		rlog.Error("Invalid setting ID format", "error", err, "id", id)
		return nil, errs.WrapCode(err, errs.InvalidArgument, "invalid_setting_id")
	}

	settingEntity, err := qrs.GetSettingByID(ctx, settingID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			rlog.Info("Setting not found", "id", id)
			return nil, errs.WrapCode(err, errs.NotFound, "setting_not_found")
		}

		rlog.Error(
			"Encountered unexpected error while querying settings table by ID",
			"error", err,
			"id", id,
		)

		return nil, errs.WrapCode(err, errs.Internal, "setting_not_retrieved")
	}

	setting, err := toDomain(settingEntity)
	if err != nil {
		rlog.Error("Failed to map setting entity to domain", "error", err)
		return nil, errs.WrapCode(err, errs.Internal, "setting_mapping_failed")
	}

	return &setting, nil
}

// BySlug finds a setting by its slug
func (r *FindSettingRepository) BySlug(ctx context.Context, qrs *db.Queries, slug string) (*entities.Setting, error) {
	settingEntity, err := qrs.GetSettingBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			rlog.Info("Setting not found", "slug", slug)
			return nil, errs.WrapCode(err, errs.NotFound, "setting_not_found")
		}

		rlog.Error(
			"Encountered unexpected error while querying settings table by slug",
			"error", err,
			"slug", slug,
		)

		return nil, errs.WrapCode(err, errs.Internal, "setting_not_retrieved")
	}

	setting, err := toDomain(settingEntity)
	if err != nil {
		rlog.Error("Failed to map setting entity to domain", "error", err)
		return nil, errs.WrapCode(err, errs.Internal, "setting_mapping_failed")
	}

	return &setting, nil
}
