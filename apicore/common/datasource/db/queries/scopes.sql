-- name: InsertScope :one
INSERT INTO scopes (
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
