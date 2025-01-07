CREATE TABLE one_time_passwords (
    "id" SERIAL PRIMARY KEY,
    "account_id" UUID REFERENCES accounts("id"),
    "expires_at" TIMESTAMPTZ NOT NULL,

    "created_at" TIMESTAMPTZ NOT NULL,
    "created_by" UUID NOT NULL,
    "deleted_at" TIMESTAMPTZ DEFAULT NULL,
    "deleted_by" UUID DEFAULT NULL,
    "updated_at" TIMESTAMPTZ DEFAULT NULL,
    "updated_by" UUID DEFAULT NULL,

    "version" INT NOT NULL
);
