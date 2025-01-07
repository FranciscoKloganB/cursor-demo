-- Insert default scopes with UUIDs as created_by/updated_by
INSERT INTO scopes (
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
        'settings-read',      -- Can read settings
        NOW(),
        '00000000-0000-0000-0000-000000000000',
        NULL,
        NULL,
        NOW(),
        '00000000-0000-0000-0000-000000000000',
        1
    ),
    (
        'settings-write',     -- Can create and update settings
        NOW(),
        '00000000-0000-0000-0000-000000000000',
        NULL,
        NULL,
        NOW(),
        '00000000-0000-0000-0000-000000000000',
        1
    ),
    (
        'settings-delete',    -- Can delete settings
        NOW(),
        '00000000-0000-0000-0000-000000000000',
        NULL,
        NULL,
        NOW(),
        '00000000-0000-0000-0000-000000000000',
        1
    );
