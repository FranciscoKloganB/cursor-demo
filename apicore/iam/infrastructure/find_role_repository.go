package infrastructure

import (
	"context"
	"database/sql"
	"errors"

	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/iam/application/ports"
	"encore.app/apicore/iam/domain/entities"
	"encore.dev/beta/errs"
	"encore.dev/rlog"
)

// FindRoleRepository implements the IFindRoleRepository interface
type FindRoleRepository struct{}

// NewFindRoleRepository creates a new instance of FindRoleRepository
func NewFindRoleRepository() ports.IFindRoleRepository {
	return &FindRoleRepository{}
}

// BySlug finds a role by its slug
func (r *FindRoleRepository) BySlug(ctx context.Context, qrs *db.Queries, slug string) (*entities.Role, error) {
	role, err := qrs.GetRoleBySlug(ctx, slug)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			rlog.Info("Role not found", "slug", slug)
			return nil, errs.WrapCode(err, errs.NotFound, "role_not_found")
		}

		rlog.Error(
			"Encountered unexpected error while querying roles table by slug",
			"error", err,
			"slug", slug,
		)

		return nil, errs.WrapCode(err, errs.Internal, "role_not_retrieved")
	}

	return &entities.Role{
		ID:        &role.ID,
		Slug:      role.Slug,
		CreatedAt: role.CreatedAt,
		CreatedBy: role.CreatedBy,
		DeletedAt: role.DeletedAt,
		DeletedBy: role.DeletedBy,
		UpdatedAt: role.UpdatedAt,
		UpdatedBy: role.UpdatedBy,
		Version:   role.Version,
	}, nil
}
