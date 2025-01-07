-- name: InsertSetting :one
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
) RETURNING *;

-- name: GetSettingByID :one
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
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetSettingBySlug :one
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
WHERE slug = $1 AND deleted_at IS NULL;
