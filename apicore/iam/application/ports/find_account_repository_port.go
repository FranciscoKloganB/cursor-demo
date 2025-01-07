package ports

import (
	"context"

	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/iam/domain/entities"
)

// IFindAccountRepository defines methods for finding an account basic information
type IFindAccountRepository interface {
	// ByEmail finds an account by the provided email
	ByEmail(ctx context.Context, qrs *db.Queries, email string) (*entities.Account, error)
}
