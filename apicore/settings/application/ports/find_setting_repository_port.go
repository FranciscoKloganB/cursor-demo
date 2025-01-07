package ports

import (
	"context"

	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/settings/domain/entities"
)

// IFindSettingRepository defines the interface for finding settings
type IFindSettingRepository interface {
	ByID(ctx context.Context, qrs *db.Queries, id string) (*entities.Setting, error)
	BySlug(ctx context.Context, qrs *db.Queries, slug string) (*entities.Setting, error)
}
