-- name: InsertOneTimePassword :one
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
) RETURNING *;
