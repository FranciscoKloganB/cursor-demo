
-- name: InsertRole :one
INSERT INTO roles (
    "slug",
    "created_at",
    "created_by",
    "deleted_at",
    "deleted_by",
    "updated_at",
    "updated_by",
    "version"
) VALUES (
    $1, $2, $3, NULL, NULL, NULL, NULL, $4
) RETURNING *;

-- name: GetRoleBySlug :one
SELECT * FROM roles
WHERE "slug" = $1 AND "deleted_at" IS NULL
LIMIT 1;
