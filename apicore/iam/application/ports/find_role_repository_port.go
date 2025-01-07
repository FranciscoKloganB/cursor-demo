package ports

import (
	"context"

	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/iam/domain/entities"
)

// IFindRoleRepository defines the interface for role retrieval operations
type IFindRoleRepository interface {
	BySlug(ctx context.Context, qrs *db.Queries, slug string) (*entities.Role, error)
}
