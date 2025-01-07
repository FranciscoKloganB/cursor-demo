-- Remove default scopes
DELETE FROM scopes
WHERE slug IN (
    'settings-read',
    'settings-write',
    'settings-delete'
);
