// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: settings.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const getSettingByID = `-- name: GetSettingByID :one
SELECT
    id,
    name,
    slug,
    hint,
    is_active,
    created_at,
    created_by,
    deleted_at,
    deleted_by,
    updated_at,
    updated_by,
    version
FROM settings
WHERE id = $1 AND deleted_at IS NULL
`

func (q *Queries) GetSettingByID(ctx context.Context, id uuid.UUID) (Setting, error) {
	row := q.db.QueryRow(ctx, getSettingByID, id)
	var i Setting
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.Hint,
		&i.IsActive,
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

const getSettingBySlug = `-- name: GetSettingBySlug :one
SELECT
    id,
    name,
    slug,
    hint,
    is_active,
    created_at,
    created_by,
    deleted_at,
    deleted_by,
    updated_at,
    updated_by,
    version
FROM settings
WHERE slug = $1 AND deleted_at IS NULL
`

func (q *Queries) GetSettingBySlug(ctx context.Context, slug string) (Setting, error) {
	row := q.db.QueryRow(ctx, getSettingBySlug, slug)
	var i Setting
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.Hint,
		&i.IsActive,
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

const insertSetting = `-- name: InsertSetting :one
INSERT INTO settings (
    "id",
    "name",
    "is_active",
    "hint",
    "slug",
    "created_at",
    "created_by",
    "deleted_at",
    "deleted_by",
    "updated_at",
    "updated_by",
    "version"
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, NULL, NULL, NULL, NULL, $8
) RETURNING id, name, slug, hint, is_active, created_at, created_by, deleted_at, deleted_by, updated_at, updated_by, version
`

type InsertSettingParams struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	IsActive  bool      `db:"is_active" json:"is_active"`
	Hint      string    `db:"hint" json:"hint"`
	Slug      string    `db:"slug" json:"slug"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	CreatedBy uuid.UUID `db:"created_by" json:"created_by"`
	Version   int32     `db:"version" json:"version"`
}

func (q *Queries) InsertSetting(ctx context.Context, arg InsertSettingParams) (Setting, error) {
	row := q.db.QueryRow(ctx, insertSetting,
		arg.ID,
		arg.Name,
		arg.IsActive,
		arg.Hint,
		arg.Slug,
		arg.CreatedAt,
		arg.CreatedBy,
		arg.Version,
	)
	var i Setting
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.Hint,
		&i.IsActive,
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
