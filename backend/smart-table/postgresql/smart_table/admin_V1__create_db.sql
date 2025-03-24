BEGIN;

CREATE SCHEMA IF NOT EXISTS smart_table_admin;

CREATE TABLE IF NOT EXISTS smart_table_admin.users (
    "uuid" UUID PRIMARY KEY NOT NULL,
    "login" TEXT NOT NULL CHECK ("login" ~ '^[a-zA-Z][a-zA-Z0-9_]{4,31}$'),
    "tg_id" TEXT NOT NULL,
    "tg_login" TEXT NOT NULL CHECK ("tg_login" ~ '^[a-zA-Z][a-zA-Z0-9_]{4,31}$'),
    "chat_id" TEXT NOT NULL,
    "first_name" TEXT NOT NULL CHECK ("first_name" ~ '^[A-Za-z\\-\\s]+$'),
    "last_name" TEXT NOT NULL CHECK ("last_name" ~ '^[A-Za-z\\-\\s]+$'),
    "password_hash" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE OR REPLACE FUNCTION set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DO $$
DECLARE tbl TEXT;
BEGIN
FOR tbl IN SELECT tablename FROM pg_tables WHERE schemaname = 'smart_table_admin' LOOP
        EXECUTE format(
            'CREATE TRIGGER trigger_%s_updated_at
            BEFORE UPDATE ON %s
            FOR EACH ROW
            EXECUTE FUNCTION set_timestamp();', tbl, tbl
        );
END LOOP;
END $$;

COMMIT;