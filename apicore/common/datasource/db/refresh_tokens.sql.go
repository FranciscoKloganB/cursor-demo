// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: refresh_tokens.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const findRefreshTokenByAccountID = `-- name: FindRefreshTokenByAccountID :one
SELECT
    "id",
    "account_id",
    "is_revoked",
    "token_value",
    "expires_at",
    "last_used_at",
    "created_at",
    "created_by",
    "deleted_at",
    "deleted_by",
    "updated_at",
    "updated_by",
    "version"
FROM
    refresh_tokens
WHERE
    "account_id" = $1
LIMIT 1
`

func (q *Queries) FindRefreshTokenByAccountID(ctx context.Context, accountID *uuid.UUID) (RefreshToken, error) {
	row := q.db.QueryRow(ctx, findRefreshTokenByAccountID, accountID)
	var i RefreshToken
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.IsRevoked,
		&i.TokenValue,
		&i.ExpiresAt,
		&i.LastUsedAt,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.Version,
	)
	return i, err
}

const findRefreshTokenByValue = `-- name: FindRefreshTokenByValue :one
SELECT
    "id",
    "account_id",
    "is_revoked",
    "token_value",
    "expires_at",
    "last_used_at",
    "created_at",
    "created_by",
    "deleted_at",
    "deleted_by",
    "updated_at",
    "updated_by",
    "version"
FROM
    refresh_tokens
WHERE
    "token_value" = $1
LIMIT 1
`

func (q *Queries) FindRefreshTokenByValue(ctx context.Context, tokenValue string) (RefreshToken, error) {
	row := q.db.QueryRow(ctx, findRefreshTokenByValue, tokenValue)
	var i RefreshToken
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.IsRevoked,
		&i.TokenValue,
		&i.ExpiresAt,
		&i.LastUsedAt,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.Version,
	)
	return i, err
}

const insertRefreshToken = `-- name: InsertRefreshToken :one
INSERT INTO refresh_tokens (
    "id", -- the refresh token UID
    "account_id", -- the account (subject UID) that owns this refresh token
    "is_revoked",
    "token_value",
    "expires_at",
    "last_used_at",
    "created_at",
    "created_by",
    "deleted_at",
    "deleted_by",
    "updated_at",
    "updated_by",
    "version"
) VALUES (
    $1, $2, FALSE, $3, $4, NULL, $5, $6, NULL, NULL, NULL, NULL, $7
) RETURNING "token_value"
`

type InsertRefreshTokenParams struct {
	ID         uuid.UUID  `db:"id" json:"id"`
	AccountID  *uuid.UUID `db:"account_id" json:"account_id"`
	TokenValue string     `db:"token_value" json:"token_value"`
	ExpiresAt  time.Time  `db:"expires_at" json:"expires_at"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	CreatedBy  uuid.UUID  `db:"created_by" json:"created_by"`
	Version    int32      `db:"version" json:"version"`
}

func (q *Queries) InsertRefreshToken(ctx context.Context, arg InsertRefreshTokenParams) (string, error) {
	row := q.db.QueryRow(ctx, insertRefreshToken,
		arg.ID,
		arg.AccountID,
		arg.TokenValue,
		arg.ExpiresAt,
		arg.CreatedAt,
		arg.CreatedBy,
		arg.Version,
	)
	var token_value string
	err := row.Scan(&token_value)
	return token_value, err
}
