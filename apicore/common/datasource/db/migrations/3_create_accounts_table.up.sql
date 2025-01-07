CREATE TABLE accounts (
    "id" UUID PRIMARY KEY,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "password_hash" VARCHAR(255) NOT NULL,
    "verification_status" VARCHAR(32) NOT NULL,

    "created_at" TIMESTAMPTZ NOT NULL,
    "created_by" UUID NOT NULL,
    "deleted_at" TIMESTAMPTZ DEFAULT NULL,
    "deleted_by" UUID DEFAULT NULL,
    "updated_at" TIMESTAMPTZ DEFAULT NULL,
    "updated_by" UUID DEFAULT NULL,

    "version" INT NOT NULL
);
