CREATE TABLE account_organisation_roles (
    "account_id" UUID REFERENCES accounts("id"),
    "organisation_id" UUID REFERENCES organisations("id"),
    "role_id" INT REFERENCES roles("id") NOT NULL,

    "created_at" TIMESTAMPTZ NOT NULL,
    "created_by" UUID NOT NULL,
    "deleted_at" TIMESTAMPTZ DEFAULT NULL,
    "deleted_by" UUID DEFAULT NULL,
    "updated_at" TIMESTAMPTZ DEFAULT NULL,
    "updated_by" UUID DEFAULT NULL,

    "version" INT NOT NULL,

    -- Primary key changed to enforce one role per account per organisation
    PRIMARY KEY ("account_id", "organisation_id")
);

-- Index for querying by organisation and role
CREATE INDEX "idx_account_organisation_roles_org"
ON account_organisation_roles("organisation_id");
