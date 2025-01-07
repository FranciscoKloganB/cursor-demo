-- name: InsertAccount :one
INSERT INTO accounts (
    "id",
    "email",
    "password_hash",
    "verification_status",
    "created_at",
    "created_by",
    "deleted_at",
    "deleted_by",
    "updated_at",
    "updated_by",
    "version"
) VALUES (
    $1, $2, $3, $4, $5, $6, NULL, NULL, NULL, NULL, $7
) RETURNING *;

-- name: FindAccountByEmail :one
SELECT * FROM accounts
WHERE "email" = $1 AND "deleted_at" IS NULL
LIMIT 1;

-- name: FindAccountByID :one
SELECT * FROM accounts
WHERE id = $1 AND "deleted_at" IS NULL
LIMIT 1;
