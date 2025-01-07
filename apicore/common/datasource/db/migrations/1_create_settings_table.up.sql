CREATE TABLE settings (
    "id" UUID PRIMARY KEY,

    "name" VARCHAR(128) NOT NULL,
    "slug" VARCHAR(128) NOT NULL,
    "hint" VARCHAR(256) NOT NULL DEFAULT '',
    "is_active" BOOLEAN NOT NULL DEFAULT FALSE,

    "created_at" TIMESTAMPTZ NOT NULL,
    "created_by" UUID NOT NULL,
    "deleted_at" TIMESTAMPTZ DEFAULT NULL,
    "deleted_by" UUID DEFAULT NULL,
    "updated_at" TIMESTAMPTZ DEFAULT NULL,
    "updated_by" UUID DEFAULT NULL,

    "version" INT NOT NULL
);
