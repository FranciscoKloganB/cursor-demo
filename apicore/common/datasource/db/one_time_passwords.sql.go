// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: one_time_passwords.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const insertOneTimePassword = `-- name: InsertOneTimePassword :one
INSERT INTO one_time_passwords (
    "account_id",
    "expires_at",
    "created_at",
    "created_by",
    "deleted_at",
    "deleted_by",
    "updated_at",
    "updated_by",
    "version"
) VALUES (
    $1, $2, $3, $4, NULL, NULL, NULL, NULL, $5
) RETURNING id, account_id, expires_at, created_at, created_by, deleted_at, deleted_by, updated_at, updated_by, version
`

type InsertOneTimePasswordParams struct {
	AccountID *uuid.UUID `db:"account_id" json:"account_id"`
	ExpiresAt time.Time  `db:"expires_at" json:"expires_at"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	CreatedBy uuid.UUID  `db:"created_by" json:"created_by"`
	Version   int32      `db:"version" json:"version"`
}

func (q *Queries) InsertOneTimePassword(ctx context.Context, arg InsertOneTimePasswordParams) (OneTimePassword, error) {
	row := q.db.QueryRow(ctx, insertOneTimePassword,
		arg.AccountID,
		arg.ExpiresAt,
		arg.CreatedAt,
		arg.CreatedBy,
		arg.Version,
	)
	var i OneTimePassword
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.ExpiresAt,
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
