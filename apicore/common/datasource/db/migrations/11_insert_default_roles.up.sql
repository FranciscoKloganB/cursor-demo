-- Insert default roles with UUIDs as created_by/updated_by
INSERT INTO roles (
    "slug",
    "created_at",
    "created_by",
    "deleted_at",
    "deleted_by",
    "updated_at",
    "updated_by",
    "version"
) VALUES
    (
        'owner',              -- Full control over organisation
        NOW(),
        '00000000-0000-0000-0000-000000000000',
        NULL,
        NULL,
        NOW(),
        '00000000-0000-0000-0000-000000000000',
        1
    ),
    (
        'administrator',      -- Full control over organisation with few exceptions (e.g., can not delete the organisation)
        NOW(),
        '00000000-0000-0000-0000-000000000000',
        NULL,
        NULL,
        NOW(),
        '00000000-0000-0000-0000-000000000000',
        1
    ),
    (
        'project-manager',    -- Can manage projects, their configurations and environments. Can not manage billing.
        NOW(),
        '00000000-0000-0000-0000-000000000000',
        NULL,
        NULL,
        NOW(),
        '00000000-0000-0000-0000-000000000000',
        1
    ),
    (
        'billing-manager',    -- Can manage billing and other financial tasks. Can not manage projects.
        NOW(),
        '00000000-0000-0000-0000-000000000000',
        NULL,
        NULL,
        NOW(),
        '00000000-0000-0000-0000-000000000000',
        1
    ),
    (
        'contributor',        -- Can contribute to projects by creating and managing settings (e.g., perfect for non-leading developers and designers)
        NOW(),
        '00000000-0000-0000-0000-000000000000',
        NULL,
        NULL,
        NOW(),
        '00000000-0000-0000-0000-000000000000',
        1
    );
