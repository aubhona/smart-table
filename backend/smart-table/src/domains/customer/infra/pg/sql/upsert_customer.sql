-- name: UpsertCustomer :one
--
-- args:
-- $1 - JSONB

INSERT INTO smart_table_customer.customers (
    uuid,
    tg_id,
    tg_login,
    avatar_link,
    chat_id,
    created_at,
    updated_at
)
SELECT
    input.uuid,
    input.tg_id,
    input.tg_login,
    input.avatar_link,
    input.chat_id,
    input.created_at,
    input.updated_at
FROM jsonb_to_record($1::jsonb) AS input(
   uuid         UUID,
   tg_id        TEXT,
   tg_login     TEXT,
   avatar_link  TEXT,
   chat_id      TEXT,
   created_at   TIMESTAMPTZ,
   updated_at   TIMESTAMPTZ
) ON CONFLICT (uuid) DO UPDATE
SET
    chat_id = EXCLUDED.chat_id,
    tg_login = EXCLUDED.tg_login
RETURNING uuid;
