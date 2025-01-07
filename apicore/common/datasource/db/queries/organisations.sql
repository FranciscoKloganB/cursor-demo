-- name: InsertOrganisation :one
INSERT INTO organisations (
    "id",
    "name",
    "created_at",
    "created_by",
    "deleted_at",
    "deleted_by",
    "updated_at",
    "updated_by",
    "version"
) VALUES (
    $1, $2, $3, $4, NULL, NULL, NULL, NULL, $5
) RETURNING *;
