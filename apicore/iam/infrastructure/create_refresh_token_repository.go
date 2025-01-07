package infrastructure

import (
	"context"

	"encore.app/apicore/common/datasource/db"
	"encore.app/apicore/iam/application/ports"
	"encore.app/apicore/iam/domain/entities"
	"encore.dev/beta/errs"
	"encore.dev/rlog"
)

// CreateRefreshTokenRepository implements the ICreateRefreshTokenRepository interface
type CreateRefreshTokenRepository struct{}

// NewCreateRefreshTokenRepository creates a new instance of CreateRefreshTokenRepository
func NewCreateRefreshTokenRepository() ports.ICreateRefreshTokenRepository {
	return &CreateRefreshTokenRepository{}
}

// Save persists a new refresh token to the database
func (r *CreateRefreshTokenRepository) Save(ctx context.Context, qrs *db.Queries, token entities.RefreshToken) (string, error) {
	tokenValue, err := qrs.InsertRefreshToken(ctx, db.InsertRefreshTokenParams{
		ID:         token.ID,
		AccountID:  &token.AccountID,
		TokenValue: token.Value,
		ExpiresAt:  token.ExpiresAt,
		CreatedAt:  token.CreatedAt,
		CreatedBy:  token.AccountID,
		Version:    1,
	})

	if err != nil {
		rlog.Error("Failed to persist refresh token",
			"error", err,
			"subject", token.AccountID,
			"token_id", token.ID,
		)

		return "", errs.WrapCode(err, errs.Internal, "refresh_token_not_saved")
	}

	rlog.Debug("Refresh token persisted successfully",
		"subject", token.AccountID,
		"token_id", token.ID,
	)

	return tokenValue, nil
}
