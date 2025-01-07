CREATE TABLE account_organisation_scopes (
    "account_id" UUID REFERENCES accounts("id"),
    "organisation_id" UUID REFERENCES organisations("id"),
    "scope_id" INT REFERENCES scopes("id"),

    "created_at" TIMESTAMPTZ NOT NULL,
    "created_by" UUID NOT NULL,
    "deleted_at" TIMESTAMPTZ DEFAULT NULL,
    "deleted_by" UUID DEFAULT NULL,
    "updated_at" TIMESTAMPTZ DEFAULT NULL,
    "updated_by" UUID DEFAULT NULL,

    "version" INT NOT NULL,

    PRIMARY KEY ("account_id", "organisation_id", "scope_id")
);

-- Compound index for querying by organisation and scope
CREATE INDEX "idx_account_organisation_scopes_org_scope"
ON account_organisation_scopes("organisation_id", "scope_id");
