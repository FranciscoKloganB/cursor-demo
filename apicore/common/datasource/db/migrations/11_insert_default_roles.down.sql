-- Remove default roles
DELETE FROM roles
WHERE slug IN (
    'owner',
    'administrator',
    'project-manager',
    'billing-manager',
    'contributor'
);
