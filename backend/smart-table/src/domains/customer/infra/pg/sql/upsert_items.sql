-- name: UpsertItems :many
--
-- args:
-- $1 - JSONB

INSERT INTO "smart-table.items" (
    uuid,
    order_uuid,
    comment,
    status,
    resolution,
    name,
    description,
    picture_link,
    weight,
    category,
    price,
    customer_uuid,
    is_draft,
    dish_uuid,
    created_at,
    updated_at
)
SELECT
    input.uuid,
    input.order_uuid,
    input.comment,
    input.status,
    input.resolution,
    input.name,
    input.description,
    input.picture_link,
    input.weight,
    input.category,
    input.price,
    input.customer_uuid,
    input.is_draft,
    input.dish_uuid,
    input.created_at,
    input.updated_at
FROM jsonb_to_recordset($1::jsonb) AS input(
   uuid          UUID,
   order_uuid    UUID,
   comment       TEXT,
   status        TEXT,
   resolution    TEXT,
   name          TEXT,
   description   TEXT,
   picture_link  TEXT,
   weight        INT,
   category      TEXT,
   price         DECIMAL,
   customer_uuid UUID,
   is_draft      BOOLEAN,
   dish_uuid     UUID,
   created_at    TIMESTAMPTZ,
   updated_at    TIMESTAMPTZ
)
ON CONFLICT (uuid) DO UPDATE
SET
    order_uuid = EXCLUDED.order_uuid,
    comment = EXCLUDED.comment,
    status = EXCLUDED.status,
    resolution = EXCLUDED.resolution,
    is_draft = EXCLUDED.is_draft
RETURNING uuid;
