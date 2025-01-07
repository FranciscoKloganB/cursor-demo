package ports

import (
	"context"

	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/settings/domain/entities"
)

// ICreateSettingRepository defines the interface for the setting repository.
type ICreateSettingRepository interface {
	Save(ctx context.Context, qrs *db.Queries, setting entities.Setting) (*entities.Setting, error)
}
