CREATE TABLE refresh_tokens (
    "id" UUID PRIMARY KEY,
    "account_id" UUID REFERENCES accounts("id"),
    "is_revoked" BOOLEAN NOT NULL DEFAULT FALSE,
    "token_value" VARCHAR(256) NOT NULL,

    "expires_at" TIMESTAMPTZ NOT NULL,
    "last_used_at" TIMESTAMPTZ DEFAULT NULL,

    "created_at" TIMESTAMPTZ NOT NULL,
    "created_by" UUID NOT NULL,
    "deleted_at" TIMESTAMPTZ DEFAULT NULL,
    "deleted_by" UUID DEFAULT NULL,
    "updated_at" TIMESTAMPTZ DEFAULT NULL,
    "updated_by" UUID DEFAULT NULL,

    "version" INT NOT NULL
);

-- Create index for querying refresh tokens table by account id
CREATE INDEX "idx_refresh_tokens_account_id"
ON refresh_tokens("account_id");

-- Create index for querying refresh tokens table by token value
CREATE INDEX "idx_refresh_tokens_token_value"
ON refresh_tokens("token_value");
