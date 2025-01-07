package ports

import (
	"context"

	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/iam/domain/entities"
)

// ICreateAccountRepository defines the interface for account creation operations.
type ICreateAccountRepository interface {
	Save(ctx context.Context, qrs *db.Queries, account entities.Account) (*entities.Account, error)
}
