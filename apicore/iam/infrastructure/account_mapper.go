package infrastructure

import (
	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/iam/domain/entities"
)

// toPersistence converts a domain account to a persistence entity.
func toPersistence(account entities.Account) (db.Account, error) {
	return db.Account{
		ID:                 account.ID,
		Email:              account.Email,
		PasswordHash:       *account.Password,
		VerificationStatus: account.VerificationStatus,
		CreatedAt:          account.CreatedAt,
		CreatedBy:          account.CreatedBy,
		DeletedAt:          account.DeletedAt,
		DeletedBy:          account.DeletedBy,
		UpdatedAt:          account.UpdatedAt,
		UpdatedBy:          account.UpdatedBy,
		Version:            account.Version,
	}, nil
}

// toDomain converts a persistence entity to a domain account.
func toDomain(entity db.Account) (entities.Account, error) {
	account, err := entities.NewAccount(
		entity.ID,
		entity.Email,
		entity.PasswordHash,
		entity.VerificationStatus,
		entity.CreatedAt,
		entity.CreatedBy,
		entity.DeletedAt,
		entity.DeletedBy,
		entity.UpdatedAt,
		entity.UpdatedBy,
		entity.Version,
	)

	return account, err
}
