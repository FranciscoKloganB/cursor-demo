-- name: InsertAccountOrganisationScope :one
INSERT INTO account_organisation_scopes (
    "account_id",
    "organisation_id",
    "scope_id",
    "created_at",
    "created_by",
    "deleted_at",
    "deleted_by",
    "updated_at",
    "updated_by",
    "version"
) VALUES (
    $1, $2, $3, $4, $5, NULL, NULL, NULL, NULL, $6
) RETURNING *;
