package ports

import (
	"context"

	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/iam/domain/entities"
)

// ICreateRefreshTokenRepository defines the interface to persist refresh tokens.
type ICreateRefreshTokenRepository interface {
	Save(ctx context.Context, qrs *db.Queries, token entities.RefreshToken) (string, error)
}
