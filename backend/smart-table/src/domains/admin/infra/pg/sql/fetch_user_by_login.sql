-- name: FetchUserByLogin :one
--
-- args:
-- $1 - TEXT

SELECT
    to_jsonb(u)
FROM
    smart_table_admin.users u
WHERE
    u.login = $1::TEXT;