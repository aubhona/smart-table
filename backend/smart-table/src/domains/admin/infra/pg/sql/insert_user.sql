-- name: InsertUser :one
--
-- args:
-- $1 - JSONB

INSERT INTO smart_table_admin.users (
    uuid,
    login,
    tg_id,
    tg_login,
    chat_id,
    first_name,
    last_name,
    password_hash,
    created_at,
    updated_at
)
SELECT
    input.uuid,
    input.login,
    input.tg_id,
    input.tg_login,
    input.chat_id,
    input.first_name,
    input.last_name,
    input.password_hash,
    input.created_at,
    input.updated_at
FROM jsonb_to_record($1::jsonb) AS input(
  uuid            UUID,
  login           TEXT,
  tg_id           TEXT,
  tg_login        TEXT,
  chat_id         TEXT,
  first_name      TEXT,
  last_name       TEXT,
  password_hash   TEXT,
  created_at      TIMESTAMPTZ,
  updated_at      TIMESTAMPTZ
)
RETURNING uuid;
