-- name: InsertAccountOrganisationRole :exec
INSERT INTO account_organisation_roles (
    "account_id",
    "organisation_id",
    "role_id",
    "created_at",
    "created_by",
    "deleted_at",
    "deleted_by",
    "updated_at",
    "updated_by",
    "version"
) VALUES (
    $1, $2, $3, $4, $5, NULL, NULL, NULL, NULL, $6
) ;

-- name: FindAccountOrganisationsRoles :many
SELECT
    -- Account data
    a.id,
    a.email,
    a.password_hash,
    a.verification_status,
    a.created_at,
    a.created_by,
    a.deleted_at,
    a.deleted_by,
    a.updated_at,
    a.updated_by,
    a.version,
    -- Organisation role data
    aor.organisation_id,
    aor.role_id,
    r.slug as role_slug
FROM
    accounts a
    LEFT JOIN account_organisation_roles aor ON a.id = aor.account_id AND aor.deleted_at IS NULL
    LEFT JOIN roles r ON r.id = aor.role_id AND r.deleted_at IS NULL
WHERE
    a.id = $1
    AND a.deleted_at IS NULL;
