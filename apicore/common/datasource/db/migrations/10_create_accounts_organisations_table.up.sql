CREATE TABLE accounts_organisations (
    "account_id" UUID REFERENCES accounts("id"),
    "organisation_id" UUID REFERENCES organisations("id"),

    "created_at" TIMESTAMPTZ NOT NULL,
    "created_by" UUID NOT NULL,
    "deleted_at" TIMESTAMPTZ DEFAULT NULL,
    "deleted_by" UUID DEFAULT NULL,
    "updated_at" TIMESTAMPTZ DEFAULT NULL,
    "updated_by" UUID DEFAULT NULL,

    "version" INT NOT NULL,

    -- Primary key creates an implicit B-tree index that can be used for:
    -- 1. Queries using both account_id AND organisation_id (filtering order matters)
    -- 2. Queries using just account_id (leftmost column principle)
    PRIMARY KEY ("account_id", "organisation_id")
);

-- Additional index for organisation_id queries since they're not covered
-- by the primary key's implicit index
CREATE INDEX "idx_accounts_organisations_organisation_id"
ON accounts_organisations("organisation_id");
