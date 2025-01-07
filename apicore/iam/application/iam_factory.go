package application

import (
	"fmt"
	"time"

	"encore.app/apicore/iam/application/commands"
	"encore.app/apicore/iam/domain/entities"
	"github.com/google/uuid"
)

// createAccountFactory creates a new Account from the given struct
//
// When the input struct cannot be converted to an Account domain object, it will return an error.
func createAccountFactory(value interface{}) (entities.Account, error) {
	switch v := value.(type) {
	case commands.CreateAccountCommand:
		now := time.Now().UTC()
		uid := uuid.New()

		account, err := entities.NewAccount(
			uid,
			v.Email,
			v.Password,
			"pending",
			now,
			uid,
			nil,
			&uuid.Nil,
			nil,
			&uuid.Nil,
			1,
		)

		if err != nil {
			return account, err
		}

		return account, nil
	default:
		return entities.Account{}, fmt.Errorf("unknown value type: %T cannot be converted to domain entity of type Account", v)
	}
}

// createRefreshTokenFactory creates a new RefreshToken entity from the given input.
//
// Returns an error if the input cannot be converted to a RefreshToken entity.
func createRefreshTokenFactory(value interface{}) (entities.RefreshToken, error) {
	switch cmd := value.(type) {
	case commands.CreateRefreshTokenCommand:
		return entities.NewRefreshToken(
			uuid.New(),
			cmd.AccountID,
			cmd.TokenValue,
			time.Now().Add(cmd.TimeToLive).UTC(),
			time.Now().UTC(),
			cmd.AccountID,
			nil,
			&uuid.Nil,
			nil,
			&uuid.Nil,
			1,
		)
	default:
		return entities.RefreshToken{}, fmt.Errorf("unsupported type: %T", value)
	}
}
