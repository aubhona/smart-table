-- name: FetchUserByUUID :one
--
-- args:
-- $1 - UUID

SELECT
    to_jsonb(u)
FROM
    smart_table_admin.users u
WHERE
    u.uuid = $1::UUID;
