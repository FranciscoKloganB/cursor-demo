-- name: InsertRefreshToken :one
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
) RETURNING "token_value";

-- name: FindRefreshTokenByAccountID :one
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
LIMIT 1;

-- name: FindRefreshTokenByValue :one
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
LIMIT 1;
