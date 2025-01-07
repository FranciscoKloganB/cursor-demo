package infrastructure

import (
	"time"

	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/settings/domain/entities"
	"github.com/google/uuid"
)

// ToPersistence converts a domain setting to a persistence entity.
func toPersistence(setting entities.Setting) (db.Setting, error) {
	var deletedAt *time.Time
	var deletedBy *uuid.UUID

	if setting.DeletedAt != nil {
		deletedAt = setting.DeletedAt
	}

	if setting.DeletedBy != nil {
		deletedBy = setting.DeletedBy
	}

	ent := db.Setting{
		ID:        setting.ID,
		Name:      setting.Name,
		Slug:      setting.Slug,
		Hint:      setting.Hint,
		IsActive:  setting.IsActive,
		CreatedAt: setting.CreatedAt,
		CreatedBy: setting.CreatedBy,
		DeletedAt: deletedAt,
		DeletedBy: deletedBy,
		UpdatedAt: setting.UpdatedAt,
		UpdatedBy: setting.UpdatedBy,
		Version:   setting.Version,
	}

	return ent, nil
}

// ToDomain converts a persistence entity to a domain setting.
func toDomain(entity db.Setting) (entities.Setting, error) {
	setting, err := entities.NewSetting(
		entity.ID,
		entity.Name,
		entity.Slug,
		entity.Hint,
		entity.IsActive,
		entity.CreatedAt,
		entity.CreatedBy,
		entity.DeletedAt,
		entity.DeletedBy,
		entity.UpdatedAt,
		entity.UpdatedBy,
		entity.Version,
	)

	return setting, err
}
